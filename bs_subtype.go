package hpdf

/*
#cgo LDFLAGS: -lhpdf -lpng -lz
#include "hpdf.h"
*/
import "C"

type BSSubtype int

const (
	BS_SOLID      BSSubtype = C.HPDF_BS_SOLID
	BS_DASHED     BSSubtype = C.HPDF_BS_DASHED
	BS_BEVELED    BSSubtype = C.HPDF_BS_BEVELED
	BS_INSET      BSSubtype = C.HPDF_BS_INSET
	BS_UNDERLINED BSSubtype = C.HPDF_BS_UNDERLINED
)
