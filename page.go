package hpdf

/*
#cgo LDFLAGS: -lhpdf -lpng -lz
#include "hpdf.h"
*/
import "C"

func (pdf *PDF) GetCurrentPage() *Page {
	page := C.HPDF_GetCurrentPage(pdf.doc)

	if page == nil {
		return nil
	}

	return &Page{page}
}

func (pdf *PDF) AddPage() (*Page, error) {
	page := C.HPDF_AddPage(pdf.doc)

	if page == nil {
		return nil, pdf.GetLastError()
	}

	return newPageBySrc(page), nil
}

func (pdf *PDF) InsertPage(src *Page) (*Page, error) {
	page := C.HPDF_InsertPage(pdf.doc, src.page)

	if page == nil {
		return nil, pdf.GetLastError()
	}

	return newPageBySrc(page), nil
}

type Page struct {
	page C.HPDF_Page
}

func NewPage() *Page {
	page := &Page{}
	return page
}

func newPageBySrc(src C.HPDF_Page) *Page {
	page := NewPage()
	page.page = src
	return page
}
