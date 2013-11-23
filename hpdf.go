package hpdf

/*
#cgo LDFLAGS: -lhpdf -lpng -lz
#include "hpdf.h"
*/
import "C"

type PDF struct {
	pdf C.HPDF_Doc
}

func New() (*PDF, error) {
	pdf := C.HPDF_New(nil, nil)

	if err := C.HPDF_GetError(pdf); err != C.HPDF_OK {
		return nil, NewError(err)
	}

	return &PDF{pdf}, nil
}

func (pdf *PDF) Free() {
	C.HPDF_Free(pdf.pdf)
}
