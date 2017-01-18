package hpdf

/*
#cgo LDFLAGS: -lhpdf -lpng -lz
#include "hpdf.h"
*/
import "C"
import "unsafe"

func Version() string {
	cString := C.HPDF_GetVersion()
	defer C.free(unsafe.Pointer(cString))

	return C.GoString(cString)
}
