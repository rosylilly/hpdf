package hpdf

/*
#cgo LDFLAGS: -lhpdf -lpng -lz
#include "hpdf.h"
*/
import "C"

type Box struct {
	Left   float32
	Bottom float32
	Right  float32
	Top    float32
}

func NewBoxByHPDFBox(box C.HPDF_Box) *Box {
	return &Box{
		float32(box.left),
		float32(box.bottom),
		float32(box.right),
		float32(box.top),
	}
}
