package hpdf

/*
#cgo LDFLAGS: -lhpdf -lpng -lz
#include "hpdf.h"
*/
import "C"

type PageLayout int

const (
	PAGE_LAYOUT_SINGLE           PageLayout = C.HPDF_PAGE_LAYOUT_SINGLE
	PAGE_LAYOUT_ONE_COLUMN       PageLayout = C.HPDF_PAGE_LAYOUT_ONE_COLUMN
	PAGE_LAYOUT_TWO_COLUMN_LEFT  PageLayout = C.HPDF_PAGE_LAYOUT_TWO_COLUMN_LEFT
	PAGE_LAYOUT_TWO_COLUMN_RIGHT PageLayout = C.HPDF_PAGE_LAYOUT_TWO_COLUMN_RIGHT
)
