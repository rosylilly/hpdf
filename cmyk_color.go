package hpdf

/*
#cgo LDFLAGS: -lhpdf -lpng -lz
#include "hpdf.h"
*/
import "C"

type CMYKColor struct {
	C float32
	M float32
	Y float32
	K float32
}

func cmykColorFromHPDF_CMYKColor(color C.HPDF_CMYKColor) *CMYKColor {
	return &CMYKColor{
		C: float32(color.c),
		M: float32(color.m),
		Y: float32(color.y),
		K: float32(color.k),
	}
}
