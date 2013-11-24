package hpdf

/*
#cgo LDFLAGS: -lhpdf -lpng -lz
#include "hpdf.h"
*/
import "C"

type Point struct {
	X float32
	Y float32
}

func pointFromHPDFPoint(point C.HPDF_Point) *Point {
	return &Point{
		float32(point.x),
		float32(point.y),
	}
}
