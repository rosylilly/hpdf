package hpdf

/*
#cgo LDFLAGS: -lhpdf -lpng -lz
#include "hpdf.h"
*/
import "C"

func (page *Page) DrawImage(image *Image, x, y, width, height float32) error {
	C.HPDF_Page_DrawImage(
		page.page, image.image,
		C.HPDF_REAL(x), C.HPDF_REAL(y), C.HPDF_REAL(width), C.HPDF_REAL(height),
	)
	return page.pdf.GetLastError()
}

func (page *Page) SetFontAndSize(font *Font, size float32) error {
	C.HPDF_Page_SetFontAndSize(
		page.page, font.font, C.HPDF_REAL(size),
	)
	return page.pdf.GetLastError()
}

func (page *Page) SetRGBFill(r float32, g float32, b float32) error {
	C.HPDF_Page_SetRGBFill(
		page.page,
		C.HPDF_REAL(r),
		C.HPDF_REAL(g),
		C.HPDF_REAL(b),
	)

	return page.pdf.GetLastError()
}

func (page *Page) SetRGBStroke(r float32, g float32, b float32) error {
	C.HPDF_Page_SetRGBStroke(
		page.page,
		C.HPDF_REAL(r),
		C.HPDF_REAL(g),
		C.HPDF_REAL(b),
	)

	return page.pdf.GetLastError()
}

func (page *Page) SetCMYKFill(c, m, y, k float32) error {
	C.HPDF_Page_SetCMYKFill(
		page.page,
		C.HPDF_REAL(c),
		C.HPDF_REAL(m),
		C.HPDF_REAL(y),
		C.HPDF_REAL(k),
	)

	return page.pdf.GetLastError()
}

func (page *Page) SetCMYKStroke(c, m, y, k float32) error {
	C.HPDF_Page_SetCMYKStroke(
		page.page,
		C.HPDF_REAL(c),
		C.HPDF_REAL(m),
		C.HPDF_REAL(y),
		C.HPDF_REAL(k),
	)

	return page.pdf.GetLastError()
}
