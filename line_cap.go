package hpdf

/*
#cgo LDFLAGS: -lhpdf -lpng -lz
#include "hpdf.h"
*/
import "C"

type LineCap int

const (
	LINE_CAP_BUTT_END              LineCap = C.HPDF_BUTT_END
	LINE_CAP_ROUND_END             LineCap = C.HPDF_ROUND_END
	LINE_CAP_PROJECTING_SCUARE_END LineCap = C.HPDF_PROJECTING_SCUARE_END
	LINE_CAP_LINECAP_EOF           LineCap = C.HPDF_LINECAP_EOF
)
