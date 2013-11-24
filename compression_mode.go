package hpdf

/*
#cgo LDFLAGS: -lhpdf -lpng -lz
#include "hpdf.h"
*/
import "C"

type CompressionMode uint32

const (
	COMP_NONE     CompressionMode = C.HPDF_COMP_NONE
	COMP_TEXT     CompressionMode = C.HPDF_COMP_TEXT
	COMP_IMAGE    CompressionMode = C.HPDF_COMP_IMAGE
	COMP_METADATA CompressionMode = C.HPDF_COMP_METADATA
	COMP_ALL      CompressionMode = C.HPDF_COMP_ALL
)
