package hpdf

/*
#cgo LDFLAGS: -lhpdf -lpng -lz
#include "hpdf.h"
*/
import "C"
import (
	"unsafe"
)

func (pdf *PDF) CreateOutline(parent *Outline, caption string, encoder *Encoder) (*Outline, error) {
	var cencoder C.HPDF_Encoder = nil
	var cparent C.HPDF_Outline = nil
	ccaption := C.CString(caption)
	if encoder != nil {
		cencoder = encoder.encoder
	}
	if parent != nil {
		cparent = parent.outline
	}
	outline := C.HPDF_CreateOutline(pdf.doc, cparent, ccaption, cencoder)
	C.free(unsafe.Pointer(ccaption))

	if outline != nil {
		return &Outline{outline, parent, caption, pdf}, nil
	} else {
		return nil, pdf.GetLastError()
	}
}

type Outline struct {
	outline C.HPDF_Outline
	Parent  *Outline
	Caption string
	pdf     *PDF
}

func (outline *Outline) SetOpened(opend bool) error {
	var copened C.HPDF_BOOL = C.HPDF_FALSE
	if opend {
		copened = C.HPDF_TRUE
	}
	C.HPDF_Outline_SetOpened(outline.outline, copened)
	return outline.pdf.GetLastError()
}

func (outline *Outline) SetDestination(destination *Destination) error {
	C.HPDF_Outline_SetDestination(outline.outline, destination.destination)
	return outline.pdf.GetLastError()
}
