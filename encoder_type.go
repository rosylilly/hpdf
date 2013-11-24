package hpdf

/*
#cgo LDFLAGS: -lhpdf -lpng -lz
#include "hpdf.h"
*/
import "C"

type EncoderType int

const (
	ENCODER_TYPE_SINGLE_BYTE   EncoderType = C.HPDF_ENCODER_TYPE_SINGLE_BYTE
	ENCODER_TYPE_DOUBLE_BYTE   EncoderType = C.HPDF_ENCODER_TYPE_DOUBLE_BYTE
	ENCODER_TYPE_UNINITIALIZED EncoderType = C.HPDF_ENCODER_TYPE_UNINITIALIZED
	ENCODER_UNKNOWN            EncoderType = C.HPDF_ENCODER_UNKNOWN
)
