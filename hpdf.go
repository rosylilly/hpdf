package hpdf

/*
#cgo LDFLAGS: -lhpdf -lpng -lz
#include "hpdf.h"
*/
import "C"
import (
	"runtime"
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

func (pdf *PDF) SaveToStream() error {
	C.HPDF_SaveToStream(pdf.doc)
	return pdf.GetLastError()
}

func (pdf *PDF) GetStreamSize() uint32 {
	return uint32(C.HPDF_GetStreamSize(pdf.doc))
}

// TODO: func (pdf *PDF) ReadFromStream(buf []byte)

func (pdf *PDF) ResetStream() error {
	C.HPDF_ResetStream(pdf.doc)
	return pdf.GetLastError()
}

func (pdf *PDF) HasDoc() (bool, error) {
	ret := C.HPDF_HasDoc(pdf.doc) == C.HPDF_TRUE
	return ret, pdf.GetLastError()
}
