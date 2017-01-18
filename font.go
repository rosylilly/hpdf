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

func (font *Font) GetFontName() string {
	cfontName := C.HPDF_Font_GetFontName(font.font)

	if cfontName != nil {
		return C.GoString(cfontName)
	} else {
		return ""
	}
}

func (font *Font) GetEncodingName() string {
	cencodingName := C.HPDF_Font_GetEncodingName(font.font)

	if cencodingName != nil {
		return C.GoString(cencodingName)
	} else {
		return ""
	}
}

func (font *Font) GetUnicodeWidth(char rune) int32 {
	ccharWidth := C.HPDF_Font_GetUnicodeWidth(font.font, C.HPDF_UNICODE(char))
	return int32(ccharWidth)
}

func (font *Font) GetBBox() *Box {
	return newBoxByHPDFBox(C.HPDF_Font_GetBBox(font.font))
}

func (font *Font) GetAscent() int32 {
	return int32(C.HPDF_Font_GetAscent(font.font))
}

func (font *Font) GetDescent() int32 {
	return int32(C.HPDF_Font_GetDescent(font.font))
}

func (font *Font) GetXHeight() uint32 {
	return uint32(C.HPDF_Font_GetXHeight(font.font))
}

func (font *Font) GetCapHeight() uint32 {
	return uint32(C.HPDF_Font_GetCapHeight(font.font))
}

func (font *Font) TextWidth(text string) *TextWidth {
	bytes := []byte(text)
	ptr := (*C.HPDF_BYTE)((unsafe.Pointer(&bytes[0])))

	ctw := C.HPDF_Font_TextWidth(font.font, ptr, C.HPDF_UINT(uint32(len(bytes))))

	tw := textWidthFromHPDF_TextWidth(ctw)

	return tw
}

func (font *Font) MeasureText(text string, width float32, fontSize float32, charSpace float32, wordSpace float32, wordwrap bool, realWidth *float32) uint32 {
	bytes := []byte(text)
	ptr := (*C.HPDF_BYTE)((unsafe.Pointer(&bytes[0])))

	var cWordwrap C.HPDF_BOOL = C.HPDF_TRUE
	if !wordwrap {
		cWordwrap = C.HPDF_FALSE
	}

	return uint32(C.HPDF_Font_MeasureText(
		font.font, ptr, C.HPDF_UINT(uint32(len(bytes))),
		C.HPDF_REAL(width), C.HPDF_REAL(fontSize), C.HPDF_REAL(charSpace), C.HPDF_REAL(wordSpace),
		cWordwrap, (*C.HPDF_REAL)(realWidth),
	))

}
