package hpdf

/*
#cgo LDFLAGS: -lhpdf -lpng -lz
#include "hpdf.h"
*/
import "C"
import (
	"fmt"
)

type Error struct {
	status C.HPDF_STATUS
}

func NewError(status C.HPDF_STATUS) *Error {
	return &Error{status}
}

func (err *Error) Error() string {
	return fmt.Sprintf("HPDF Error: 0x%04X", err.status)
}
