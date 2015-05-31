package hpdf

/*
#cgo LDFLAGS: -lhpdf -lpng -lz
#include "hpdf.h"
*/
import "C"

func Version() string {
	cString := C.HPDF_GetVersion()

	return C.GoString(cString)
}
