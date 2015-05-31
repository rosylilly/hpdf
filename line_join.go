package hpdf

/*
#cgo LDFLAGS: -lhpdf -lpng -lz
#include "hpdf.h"
*/
import "C"

type LineJoin int

const (
	LINE_JOIN_MITER_JOIN   LineJoin = C.HPDF_MITER_JOIN
	LINE_JOIN_ROUND_JOIN   LineJoin = C.HPDF_ROUND_JOIN
	LINE_JOIN_BEVEL_JOIN   LineJoin = C.HPDF_BEVEL_JOIN
	LINE_JOIN_LINEJOIN_EOF LineJoin = C.HPDF_LINEJOIN_EOF
)
