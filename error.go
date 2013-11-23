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
	Code   ErrorCode
	status C.HPDF_STATUS
	detail C.HPDF_STATUS
}

func NewError(status, detail C.HPDF_STATUS) *Error {
	return &Error{ErrorCode(status), status, detail}
}

func (err *Error) Error() string {
	return fmt.Sprintf(
		"HPDF Error: 0x%04X-0x%04X: %s",
		err.status,
		err.detail,
		err.message(),
	)
}

func (err *Error) message() string {
	msg, hit := ErrorMessages[err.Code]

	if !hit {
		return "Unknown"
	}

	return msg
}
