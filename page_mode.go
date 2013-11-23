package hpdf

/*
#cgo LDFLAGS: -lhpdf -lpng -lz
#include "hpdf.h"
*/
import "C"

type PageMode int

const (
	PAGE_MODE_USE_NONE    PageMode = C.HPDF_PAGE_MODE_USE_NONE
	PAGE_MODE_USE_OUTLINE PageMode = C.HPDF_PAGE_MODE_USE_OUTLINE
	PAGE_MODE_USE_THUMBS  PageMode = C.HPDF_PAGE_MODE_USE_THUMBS
	PAGE_MODE_FULL_SCREEN PageMode = C.HPDF_PAGE_MODE_FULL_SCREEN
)
