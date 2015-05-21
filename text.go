package hpdf

/*
#cgo LDFLAGS: -lhpdf -lpng -lz
#include "hpdf.h"
*/
import "C"

func (page *Page) BeginText() error {
	C.HPDF_Page_BeginText(page.page)
	return page.pdf.GetLastError()
}

func (page *Page) EndText() error {
	C.HPDF_Page_EndText(page.page)
	return page.pdf.GetLastError()
}

func (page *Page) ShowText(text string) error {
	C.HPDF_Page_ShowText(page.page, C.CString(text))
	return page.pdf.GetLastError()
}

func (page *Page) TextOut(x float32, y float32, text string) error {
	C.HPDF_Page_TextOut(page.page, C.HPDF_REAL(x), C.HPDF_REAL(y), C.CString(text))
	return page.pdf.GetLastError()
}
