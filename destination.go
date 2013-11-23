package hpdf

/*
#cgo LDFLAGS: -lhpdf -lpng -lz
#include "hpdf.h"
*/
import "C"

type Destination struct {
	destination C.HPDF_Destination
	page        *Page
}
