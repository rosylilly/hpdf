package hpdf

/*
#cgo LDFLAGS: -lhpdf -lpng -lz
#include "hpdf.h"
*/
import "C"
import (
	"unsafe"
)

type PDF struct {
	doc C.HPDF_Doc
}

func New() (*PDF, error) {
	doc := C.HPDF_New(nil, nil)

	if err := C.HPDF_GetError(doc); err != C.HPDF_OK {
		return nil, NewError(err, 0)
	}

	return &PDF{doc}, nil
}

func (pdf *PDF) Free() {
	C.HPDF_Free(pdf.doc)
}

func (pdf *PDF) GetLastError() error {
	if err := C.HPDF_GetError(pdf.doc); err != C.HPDF_OK {
		detail := C.HPDF_GetErrorDetail(pdf.doc)
		return NewError(err, detail)
	}
	return nil
}

func (pdf *PDF) SaveToFile(fn string) error {
	cfn := C.CString(fn)
	C.HPDF_SaveToFile(pdf.doc, cfn)
	C.free(unsafe.Pointer(cfn))

	return pdf.GetLastError()
}
