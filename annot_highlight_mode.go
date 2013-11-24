package hpdf

/*
#cgo LDFLAGS: -lhpdf -lpng -lz
#include "hpdf.h"
*/
import "C"

type AnnotHighlightMode int

const (
	ANNOT_NO_HIGHTLIGHT       AnnotHighlightMode = C.HPDF_ANNOT_NO_HIGHTLIGHT
	ANNOT_INVERT_BOX          AnnotHighlightMode = C.HPDF_ANNOT_INVERT_BOX
	ANNOT_INVERT_BORDER       AnnotHighlightMode = C.HPDF_ANNOT_INVERT_BORDER
	ANNOT_DOWN_APPEARANCE     AnnotHighlightMode = C.HPDF_ANNOT_DOWN_APPEARANCE
	ANNOT_HIGHTLIGHT_MODE_EOF AnnotHighlightMode = C.HPDF_ANNOT_HIGHTLIGHT_MODE_EOF
)
