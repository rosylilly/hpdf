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

func (page *Page) CreateDestination() (*Destination, error) {
	destination := C.HPDF_Page_CreateDestination(page.page)

	if destination != nil {
		return &Destination{destination, page}, nil
	} else {
		return nil, page.pdf.GetLastError()
	}
}
