package hpdf

/*
#cgo LDFLAGS: -lhpdf -lpng -lz
#include "hpdf.h"
*/
import "C"

type TransMatrix struct {
	A float32
	B float32
	C float32
	D float32
	X float32
	Y float32
}

func transMatrixFromHPDFTransMatrix(transMatrix C.HPDF_TransMatrix) *TransMatrix {
	return &TransMatrix{
		float32(transMatrix.a),
		float32(transMatrix.b),
		float32(transMatrix.c),
		float32(transMatrix.d),
		float32(transMatrix.x),
		float32(transMatrix.y),
	}
}
