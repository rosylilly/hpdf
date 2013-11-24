package hpdf

/*
#cgo LDFLAGS: -lhpdf -lpng -lz
#include "hpdf.h"
*/
import "C"

type Rect struct {
	Left   float32
	Bottom float32
	Right  float32
	Top    float32
}

func NewRect(left, bottom, right, top float32) *Rect {
	return &Rect{left, bottom, right, top}
}

func newRectByHPDFRect(rect C.HPDF_Rect) *Rect {
	return &Rect{
		float32(rect.left),
		float32(rect.bottom),
		float32(rect.right),
		float32(rect.top),
	}
}

func (rect *Rect) toHpdf() C.HPDF_Rect {
	var hpdfRect C.HPDF_Rect

	hpdfRect.left = C.HPDF_REAL(rect.Left)
	hpdfRect.bottom = C.HPDF_REAL(rect.Bottom)
	hpdfRect.right = C.HPDF_REAL(rect.Right)
	hpdfRect.top = C.HPDF_REAL(rect.Top)

	return hpdfRect
}
