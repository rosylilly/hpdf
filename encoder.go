package hpdf

/*
#cgo LDFLAGS: -lhpdf -lpng -lz
#include "hpdf.h"
*/
import "C"
import (
	"unsafe"
)

func (pdf *PDF) GetEncoder(encodingName string) (*Encoder, error) {
	cencodingName := C.CString(encodingName)
	encoder := C.HPDF_GetEncoder(pdf.doc, cencodingName)
	C.free(unsafe.Pointer(cencodingName))

	if encoder != nil {
		return &Encoder{encoder}, nil
	} else {
		return nil, pdf.GetLastError()
	}
}

func (pdf *PDF) GetCurrentEncoder() *Encoder {
	encoder := C.HPDF_GetCurrentEncoder(pdf.doc)

	if encoder != nil {
		return &Encoder{encoder}
	} else {
		return nil
	}
}

func (pdf *PDF) SetCurrentEncoder(encodingName string) error {
	cencodingName := C.CString(encodingName)
	C.HPDF_SetCurrentEncoder(pdf.doc, cencodingName)
	C.free(unsafe.Pointer(cencodingName))

	return pdf.GetLastError()
}

type Encoder struct {
	encoding C.HPDF_Encoder
}
