package hpdf

/*
#cgo LDFLAGS: -lhpdf -lpng -lz
#include "hpdf.h"
*/
import "C"

type ViewerPreference uint32

const (
	HIDE_TOOLBAR       ViewerPreference = C.HPDF_HIDE_TOOLBAR
	HIDE_MENUBAR       ViewerPreference = C.HPDF_HIDE_MENUBAR
	HIDE_WINDOW_UI     ViewerPreference = C.HPDF_HIDE_WINDOW_UI
	FIT_WINDOW         ViewerPreference = C.HPDF_FIT_WINDOW
	CENTER_WINDOW      ViewerPreference = C.HPDF_CENTER_WINDOW
	PRINT_SCALING_NONE ViewerPreference = C.HPDF_PRINT_SCALING_NONE
)
