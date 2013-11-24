package hpdf

/*
#cgo LDFLAGS: -lhpdf -lpng -lz
#include "hpdf.h"
*/
import "C"

type WritingMode int

const (
	WMODE_HORIZONTAL WritingMode = C.HPDF_WMODE_HORIZONTAL
	WMODE_VERTICAL   WritingMode = C.HPDF_WMODE_VERTICAL
)
