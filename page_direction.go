package hpdf

/*
#cgo LDFLAGS: -lhpdf -lpng -lz
#include "hpdf.h"
*/
import "C"

type PageDirection int

const (
	PAGE_PORTRAIT  PageDirection = C.HPDF_PAGE_PORTRAIT
	PAGE_LANDSCAPE PageDirection = C.HPDF_PAGE_LANDSCAPE
)
