package hpdf

/*
#cgo LDFLAGS: -lhpdf -lpng -lz
#include "hpdf.h"
*/
import "C"

type AnnotIcon int

const (
	ANNOT_ICON_COMMENT       AnnotIcon = C.HPDF_ANNOT_ICON_COMMENT
	ANNOT_ICON_KEY           AnnotIcon = C.HPDF_ANNOT_ICON_KEY
	ANNOT_ICON_NOTE          AnnotIcon = C.HPDF_ANNOT_ICON_NOTE
	ANNOT_ICON_HELP          AnnotIcon = C.HPDF_ANNOT_ICON_HELP
	ANNOT_ICON_NEW_PARAGRAPH AnnotIcon = C.HPDF_ANNOT_ICON_NEW_PARAGRAPH
	ANNOT_ICON_PARAGRAPH     AnnotIcon = C.HPDF_ANNOT_ICON_PARAGRAPH
	ANNOT_ICON_INSERT        AnnotIcon = C.HPDF_ANNOT_ICON_INSERT
)
