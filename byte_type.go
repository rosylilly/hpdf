package hpdf

/*
#cgo LDFLAGS: -lhpdf -lpng -lz
#include "hpdf.h"
*/
import "C"

type ByteType int

const (
	BYTE_TYPE_SINGLE  ByteType = C.HPDF_BYTE_TYPE_SINGLE
	BYTE_TYPE_LEAD    ByteType = C.HPDF_BYTE_TYPE_LEAD
	BYTE_TYPE_TRIAL   ByteType = C.HPDF_BYTE_TYPE_TRIAL
	BYTE_TYPE_UNKNOWN ByteType = C.HPDF_BYTE_TYPE_UNKNOWN
)
