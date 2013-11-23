package hpdf

/*
#cgo LDFLAGS: -lhpdf -lpng -lz
#include "hpdf.h"
*/
import "C"
import (
	"unsafe"
)

func (pdf *PDF) GetFont(fontName string, encodingName ...string) (*Font, error) {
	var cencodingName *C.char = nil

	cfontName := C.CString(fontName)
	if len(encodingName) > 0 {
		cencodingName = C.CString(encodingName[0])
	}
	font := C.HPDF_GetFont(pdf.doc, cfontName, cencodingName)
	C.free(unsafe.Pointer(cfontName))
	if cencodingName != nil {
		C.free(unsafe.Pointer(cencodingName))
	}

	if font == nil {
		return nil, pdf.GetLastError()
	}

	return newFont(font), pdf.GetLastError()
}

func (pdf *PDF) LoadType1FontFromFile(afmFn string, pfmFn ...string) (string, error) {
	var cpfmFn *C.char = nil

	cafmFn := C.CString(afmFn)
	if len(pfmFn) > 0 {
		cpfmFn = C.CString(pfmFn[0])
	}
	fontName := C.HPDF_LoadType1FontFromFile(pdf.doc, cafmFn, cpfmFn)
	C.free(unsafe.Pointer(cafmFn))
	if cpfmFn != nil {
		C.free(unsafe.Pointer(cpfmFn))
	}

	if fontName != nil {
		defer C.free(unsafe.Pointer(fontName))
		return C.GoString(fontName), nil
	} else {
		return "", pdf.GetLastError()
	}
}

func (pdf *PDF) LoadTTFontFromFile(fontName string, embedding bool) (string, error) {
	cfontName := C.CString(fontName)
	var cembedding C.HPDF_BOOL = C.HPDF_TRUE
	if !embedding {
		cembedding = C.HPDF_FALSE
	}
	retFontName := C.HPDF_LoadTTFontFromFile(pdf.doc, cfontName, cembedding)
	C.free(unsafe.Pointer(cfontName))

	if retFontName != nil {
		defer C.free(unsafe.Pointer(retFontName))
		return C.GoString(retFontName), nil
	} else {
		return "", pdf.GetLastError()
	}
}

func (pdf *PDF) LoadTTFontFromFile2(fontName string, index uint, embedding bool) (string, error) {
	cfontName := C.CString(fontName)
	var cembedding C.HPDF_BOOL = C.HPDF_TRUE
	if !embedding {
		cembedding = C.HPDF_FALSE
	}
	retFontName := C.HPDF_LoadTTFontFromFile2(pdf.doc, cfontName, C.HPDF_UINT(index), cembedding)
	C.free(unsafe.Pointer(cfontName))

	if retFontName != nil {
		defer C.free(unsafe.Pointer(retFontName))
		return C.GoString(retFontName), nil
	} else {
		return "", pdf.GetLastError()
	}
}

func (pdf *PDF) UseJPFonts() error {
	C.HPDF_UseJPFonts(pdf.doc)
	return pdf.GetLastError()
}

func (pdf *PDF) UseKRFonts() error {
	C.HPDF_UseKRFonts(pdf.doc)
	return pdf.GetLastError()
}

func (pdf *PDF) UseCNSFonts() error {
	C.HPDF_UseCNSFonts(pdf.doc)
	return pdf.GetLastError()
}

func (pdf *PDF) UseCNTFonts() error {
	C.HPDF_UseCNTFonts(pdf.doc)
	return pdf.GetLastError()
}

type Font struct {
	font C.HPDF_Font
}

func newFont(src C.HPDF_Font) *Font {
	font := &Font{src}
	return font
}
