package hpdf

/*
#cgo LDFLAGS: -lhpdf -lpng -lz
#include "hpdf.h"
*/
import "C"

type RGBColor struct {
	R float32
	G float32
	B float32
}

func rgbColorFromHPDF_RGBColor(color C.HPDF_RGBColor) *RGBColor {
	return &RGBColor{
		R: float32(color.r),
		G: float32(color.g),
		B: float32(color.b),
	}
}
