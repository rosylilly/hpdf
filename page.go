package hpdf

/*
#cgo LDFLAGS: -lhpdf -lpng -lz
#include "hpdf.h"
*/
import "C"
import (
	"unsafe"
)

func (pdf *PDF) GetCurrentPage() *Page {
	page := C.HPDF_GetCurrentPage(pdf.doc)

	if page == nil {
		return nil
	}

	return &Page{page, pdf}
}

func (pdf *PDF) AddPage() (*Page, error) {
	page := C.HPDF_AddPage(pdf.doc)

	if page == nil {
		return nil, pdf.GetLastError()
	}

	return newPage(page, pdf), nil
}

func (pdf *PDF) InsertPage(src *Page) (*Page, error) {
	page := C.HPDF_InsertPage(pdf.doc, src.page)

	if page == nil {
		return nil, pdf.GetLastError()
	}

	return newPage(page, pdf), nil
}

type Page struct {
	page C.HPDF_Page
	pdf  *PDF
}

func newPage(src C.HPDF_Page, pdf *PDF) *Page {
	page := &Page{src, pdf}
	return page
}

func (page *Page) SetWidth(width float32) error {
	C.HPDF_Page_SetWidth(page.page, C.HPDF_REAL(width))
	return page.pdf.GetLastError()
}

func (page *Page) SetHeight(height float32) error {
	C.HPDF_Page_SetHeight(page.page, C.HPDF_REAL(height))
	return page.pdf.GetLastError()
}

func (page *Page) SetSize(size PageSize, direction PageDirection) error {
	C.HPDF_Page_SetSize(
		page.page, C.HPDF_PageSizes(size), C.HPDF_PageDirection(direction),
	)
	return page.pdf.GetLastError()
}

func (page *Page) SetRotate(rotate uint16) error {
	C.HPDF_Page_SetRotate(page.page, C.HPDF_UINT16(rotate))
	return page.pdf.GetLastError()
}

func (page *Page) GetWidth() float32 {
	return float32(C.HPDF_Page_GetWidth(page.page))
}

func (page *Page) GetHeight() float32 {
	return float32(C.HPDF_Page_GetHeight(page.page))
}

func (page *Page) TextWidth(text string) float32 {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	return float32(C.HPDF_Page_TextWidth(page.page, ctext))
}

func (page *Page) MeasureText(text string, width float32, wordwrap bool) (float32, error) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))

	var cwordwrap C.HPDF_BOOL = C.HPDF_FALSE
	if wordwrap {
		cwordwrap = C.HPDF_TRUE
	}

	var realWidth float32

	C.HPDF_Page_MeasureText(page.page, ctext, C.HPDF_REAL(width), cwordwrap, (*C.HPDF_REAL)(&realWidth))
	return realWidth, page.pdf.GetLastError()
}

func (page *Page) GetGMode() GMode {
	cgMode := C.HPDF_Page_GetGMode(page.page)
	return GMode(cgMode)
}

func (page *Page) GetCurrentPos() *Point {
	cpoint := C.HPDF_Page_GetCurrentPos(page.page)

	return pointFromHPDFPoint(cpoint)
}

func (page *Page) GetCurrentTextPos() *Point {
	cpoint := C.HPDF_Page_GetCurrentTextPos(page.page)

	return pointFromHPDFPoint(cpoint)
}

func (page *Page) GetCurrentFont() *Font {
	cFont := C.HPDF_Page_GetCurrentFont(page.page)

	if cFont != nil {
		return newFont(cFont)
	} else {
		return nil
	}
}

func (page *Page) GetCurrentFontSize() float32 {
	cSize := C.HPDF_Page_GetCurrentFontSize(page.page)

	return float32(cSize)
}

func (page *Page) GetTransMatrix() *TransMatrix {
	cTransMatrix := C.HPDF_Page_GetTransMatrix(page.page)

	return transMatrixFromHPDFTransMatrix(cTransMatrix)
}

func (page *Page) GetLineWidth() float32 {
	cLineWidth := C.HPDF_Page_GetLineWidth(page.page)

	return float32(cLineWidth)
}

func (page *Page) GetLineCap() LineCap {
	cLineCap := C.HPDF_Page_GetLineCap(page.page)

	return LineCap(cLineCap)
}
