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
