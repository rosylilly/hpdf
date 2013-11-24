package hpdf

/*
#cgo LDFLAGS: -lhpdf -lpng -lz
#include "hpdf.h"
*/
import "C"

type Permission uint32

const (
	ENABLE_READ     Permission = C.HPDF_ENABLE_READ
	ENABLE_PRINT    Permission = C.HPDF_ENABLE_PRINT
	ENABLE_EDIT_ALL Permission = C.HPDF_ENABLE_EDIT_ALL
	ENABLE_COPY     Permission = C.HPDF_ENABLE_COPY
	ENABLE_EDIT     Permission = C.HPDF_ENABLE_EDIT
)
