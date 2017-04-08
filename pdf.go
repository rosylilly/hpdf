package hpdf

/*
#cgo LDFLAGS: -lhpdf -lpng -lz
#include "hpdf.h"
*/
import "C"
import (
	"io"
	"runtime"
	"time"
	"unsafe"
)

type PDF struct {
	doc C.HPDF_Doc
}

func finalizePDF(pdf *PDF) {
	pdf.FreeDocAll()
	pdf.Free()
}

func New() (*PDF, error) {
	doc := C.HPDF_New(nil, nil)

	if err := C.HPDF_GetError(doc); err != C.HPDF_OK {
		return nil, NewError(err, 0)
	}

	pdf := &PDF{doc}
	runtime.SetFinalizer(pdf, finalizePDF)
	return pdf, nil
}

func (pdf *PDF) GetLastError() error {
	defer C.HPDF_ResetError(pdf.doc)
	if err := C.HPDF_GetError(pdf.doc); err != C.HPDF_OK {
		detail := C.HPDF_GetErrorDetail(pdf.doc)
		return NewError(err, detail)
	}
	return nil
}

func (pdf *PDF) Free() {
	C.HPDF_Free(pdf.doc)
}

func (pdf *PDF) NewDoc() error {
	C.HPDF_NewDoc(pdf.doc)
	return pdf.GetLastError()
}

func (pdf *PDF) FreeDoc() {
	C.HPDF_FreeDoc(pdf.doc)
}

func (pdf *PDF) FreeDocAll() {
	C.HPDF_FreeDocAll(pdf.doc)
}

func (pdf *PDF) SaveToFile(fn string) error {
	cfn := C.CString(fn)
	C.HPDF_SaveToFile(pdf.doc, cfn)
	C.free(unsafe.Pointer(cfn))

	return pdf.GetLastError()
}

func (pdf *PDF) SaveToStream(wr io.Writer) error {
	if status := C.HPDF_SaveToStream(pdf.doc); status != C.HPDF_OK {
		return pdf.GetLastError()
	}
	if status := C.HPDF_ResetStream(pdf.doc); status != C.HPDF_OK {
		return pdf.GetLastError()
	}

	const blocksize = 4096
	var buf [blocksize]byte
	ptr := (*C.HPDF_BYTE)((unsafe.Pointer(&buf[0])))
	for {
		size := C.HPDF_UINT32(blocksize)
		C.HPDF_ReadFromStream(pdf.doc, ptr, &size)
		if size == 0 {
			break
		}
		if _, err := wr.Write(buf[:size]); err != nil {
			return err
		}
	}

	return pdf.GetLastError()
}

func (pdf *PDF) HasDoc() bool {
	return C.HPDF_HasDoc(pdf.doc) == C.HPDF_TRUE
}

func (pdf *PDF) SetPagesConfiguration(page_per_pages uint) error {
	C.HPDF_SetPagesConfiguration(pdf.doc, C.HPDF_UINT(page_per_pages))
	return pdf.GetLastError()
}

func (pdf *PDF) SetPageLayout(layout PageLayout) error {
	C.HPDF_SetPageLayout(pdf.doc, C.HPDF_PageLayout(layout))
	return pdf.GetLastError()
}

func (pdf *PDF) GetPageLayout() PageLayout {
	return PageLayout(C.HPDF_GetPageLayout(pdf.doc))
}

func (pdf *PDF) SetPageMode(pageMode PageMode) error {
	C.HPDF_SetPageMode(pdf.doc, C.HPDF_PageMode(pageMode))
	return pdf.GetLastError()
}

func (pdf *PDF) GetPageMode() PageMode {
	return PageMode(C.HPDF_GetPageMode(pdf.doc))
}

func (pdf *PDF) SetOpenAction(destination *Destination) error {
	C.HPDF_SetOpenAction(pdf.doc, destination.destination)
	return pdf.GetLastError()
}

func (pdf *PDF) AddPageLabel(
	pageNum uint, pageNumStyle PageNumStyle, firstPage uint, prefix string,
) error {
	cprefix := C.CString(prefix)
	C.HPDF_AddPageLabel(
		pdf.doc,
		C.HPDF_UINT(pageNum),
		C.HPDF_PageNumStyle(pageNumStyle),
		C.HPDF_UINT(firstPage),
		cprefix,
	)
	C.free(unsafe.Pointer(cprefix))
	return nil
}

func (pdf *PDF) SetInfoAttr(infoType InfoType, value string) error {
	cvalue := C.CString(value)
	defer C.free(unsafe.Pointer(cvalue))
	C.HPDF_SetInfoAttr(pdf.doc, C.HPDF_InfoType(infoType), cvalue)
	return pdf.GetLastError()
}

func (pdf *PDF) GetInfoAttr(infoType InfoType) (string, error) {
	cvalue := C.HPDF_GetInfoAttr(pdf.doc, C.HPDF_InfoType(infoType))

	if cvalue != nil {
		defer C.free(unsafe.Pointer(cvalue))
		return C.GoString(cvalue), nil
	} else {
		return "", pdf.GetLastError()
	}
}

func (pdf *PDF) SetInfoDateAttr(infoType InfoType, datetime time.Time) error {
	cdate := timeToHPDFDate(datetime)
	C.HPDF_SetInfoDateAttr(pdf.doc, C.HPDF_InfoType(infoType), cdate)
	return pdf.GetLastError()
}

func (pdf *PDF) SetPassword(ownerPassword, userPassword string) error {
	cownerPassword := C.CString(ownerPassword)
	cuserPassword := C.CString(userPassword)
	C.HPDF_SetPassword(pdf.doc, cownerPassword, cuserPassword)
	C.free(unsafe.Pointer(cownerPassword))
	C.free(unsafe.Pointer(cuserPassword))
	return pdf.GetLastError()
}

func (pdf *PDF) SetPermission(permission Permission) error {
	C.HPDF_SetPermission(pdf.doc, C.HPDF_UINT(permission))
	return pdf.GetLastError()
}

func (pdf *PDF) GetViewerPreference() ViewerPreference {
	return ViewerPreference(C.HPDF_GetViewerPreference(pdf.doc))
}

func (pdf *PDF) SetViewerPreference(preference ViewerPreference) error {
	C.HPDF_SetViewerPreference(pdf.doc, C.HPDF_UINT(preference))
	return pdf.GetLastError()
}

func (pdf *PDF) SetEncryptMode(encryptMode EncryptMode, keyLen uint32) error {
	C.HPDF_SetEncryptionMode(
		pdf.doc, C.HPDF_EncryptMode(encryptMode), C.HPDF_UINT(keyLen),
	)
	return pdf.GetLastError()
}

func (pdf *PDF) SetCompressionMode(compressionMode CompressionMode) error {
	C.HPDF_SetCompressionMode(pdf.doc, C.HPDF_UINT(compressionMode))
	return pdf.GetLastError()
}
