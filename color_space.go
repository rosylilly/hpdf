package hpdf

/*
#cgo LDFLAGS: -lhpdf -lpng -lz
#include "hpdf.h"
*/
import "C"

type ColorSpace int

const (
	CS_DEVICE_GRAY ColorSpace = C.HPDF_CS_DEVICE_GRAY
	CS_DEVICE_RGB  ColorSpace = C.HPDF_CS_DEVICE_RGB
	CS_DEVICE_CMYK ColorSpace = C.HPDF_CS_DEVICE_CMYK
	CS_CAL_GRAY    ColorSpace = C.HPDF_CS_CAL_GRAY
	CS_CAL_RGB     ColorSpace = C.HPDF_CS_CAL_RGB
	CS_LAB         ColorSpace = C.HPDF_CS_LAB
	CS_ICC_BASED   ColorSpace = C.HPDF_CS_ICC_BASED
	CS_SEPARATION  ColorSpace = C.HPDF_CS_SEPARATION
	CS_DEVICE_N    ColorSpace = C.HPDF_CS_DEVICE_N
	CS_INDEXED     ColorSpace = C.HPDF_CS_INDEXED
	CS_PATTERN     ColorSpace = C.HPDF_CS_PATTERN
	CS_EOF         ColorSpace = C.HPDF_CS_EOF
)
