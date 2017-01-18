package hpdf

/*
#cgo LDFLAGS: -lhpdf -lpng -lz
#include "hpdf.h"
*/
import "C"

type TextWidth struct {
	NumChars uint32
	NumWords uint32
	Width    uint32
	NumSpace uint32
}

func textWidthFromHPDF_TextWidth(tw C.HPDF_TextWidth) *TextWidth {
	return &TextWidth{
		NumChars: uint32(tw.numchars),
		NumWords: uint32(tw.numwords),
		Width:    uint32(tw.width),
		NumSpace: uint32(tw.numspace),
	}
}
