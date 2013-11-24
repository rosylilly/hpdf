package hpdf

/*
#cgo LDFLAGS: -lhpdf -lpng -lz
#include "hpdf.h"
*/
import "C"

type InfoType int

const (
	INFO_CREATION_DATE InfoType = C.HPDF_INFO_CREATION_DATE
	INFO_MOD_DATE      InfoType = C.HPDF_INFO_MOD_DATE
	INFO_AUTHOR        InfoType = C.HPDF_INFO_AUTHOR
	INFO_CREATOR       InfoType = C.HPDF_INFO_CREATOR
	INFO_TITLE         InfoType = C.HPDF_INFO_TITLE
	INFO_SUBJECT       InfoType = C.HPDF_INFO_SUBJECT
	INFO_KEYWORDS      InfoType = C.HPDF_INFO_KEYWORDS
)
