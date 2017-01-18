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

func (page *Page) GetGrayFill() float32 {
	return float32(C.HPDF_Page_GetGrayFill(page.page))
}

func (page *Page) SetGrayFill(gray float32) error {
	C.HPDF_Page_SetGrayFill(page.page, C.HPDF_REAL(gray))

	return page.pdf.GetLastError()
}

func (page *Page) GetGrayStroke() float32 {
	return float32(C.HPDF_Page_GetGrayStroke(page.page))
}

func (page *Page) SetGrayStroke(gray float32) error {
	C.HPDF_Page_SetGrayStroke(page.page, C.HPDF_REAL(gray))

	return page.pdf.GetLastError()
}

func (page *Page) GetRGBFill() *RGBColor {
	return rgbColorFromHPDF_RGBColor(C.HPDF_Page_GetRGBFill(page.page))
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

func (page *Page) GetRGBStroke() *RGBColor {
	return rgbColorFromHPDF_RGBColor(C.HPDF_Page_GetRGBStroke(page.page))
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

func (page *Page) GetCMYKFill() *CMYKColor {
	return cmykColorFromHPDF_CMYKColor(C.HPDF_Page_GetCMYKFill(page.page))
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

func (page *Page) GetCMYKStroke() *CMYKColor {
	return cmykColorFromHPDF_CMYKColor(C.HPDF_Page_GetCMYKStroke(page.page))
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
