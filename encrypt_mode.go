package hpdf

/*
#cgo LDFLAGS: -lhpdf -lpng -lz
#include "hpdf.h"
*/
import "C"

type EncryptMode int

const (
	ENCRYPT_R2 EncryptMode = C.HPDF_ENCRYPT_R2
	ENCRYPT_R3 EncryptMode = C.HPDF_ENCRYPT_R3
)
