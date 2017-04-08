package hpdf

/*
#cgo LDFLAGS: -lhpdf -lpng -lz
#include "hpdf.h"
*/
import "C"

func (page *Page) CurveTo(x1 float32, y1 float32, x2 float32, y2 float32, x3 float32, y3 float32) error {
	C.HPDF_Page_CurveTo(page.page,
		C.HPDF_REAL(x1), C.HPDF_REAL(y1), C.HPDF_REAL(x2), C.HPDF_REAL(y2), C.HPDF_REAL(x3), C.HPDF_REAL(y3))

	return page.pdf.GetLastError()
}

func (page *Page) DrawImage(image *Image, x, y, width, height float32) error {
	C.HPDF_Page_DrawImage(
		page.page, image.image,
		C.HPDF_REAL(x), C.HPDF_REAL(y), C.HPDF_REAL(width), C.HPDF_REAL(height),
	)

	return page.pdf.GetLastError()
}

func (page *Page) Fill() error {
	C.HPDF_Page_Fill(page.page)

	return page.pdf.GetLastError()
}

func (page *Page) FillStroke() error {
	C.HPDF_Page_FillStroke(page.page)

	return page.pdf.GetLastError()
}

func (page *Page) LineTo(x float32, y float32) error {
	C.HPDF_Page_LineTo(page.page, C.HPDF_REAL(x), C.HPDF_REAL(y))

	return page.pdf.GetLastError()
}

func (page *Page) MoveTo(x float32, y float32) error {
	C.HPDF_Page_MoveTo(page.page, C.HPDF_REAL(x), C.HPDF_REAL(y))

	return page.pdf.GetLastError()
}

func (page *Page) Rectangle(x float32, y float32, width float32, height float32) error {
	C.HPDF_Page_Rectangle(page.page, C.HPDF_REAL(x), C.HPDF_REAL(y), C.HPDF_REAL(width), C.HPDF_REAL(height))

	return page.pdf.GetLastError()
}

func (page *Page) SetDash(dashPtn []uint16, phase uint) error {
	var ptn *C.HPDF_UINT16 = nil
	if len(dashPtn) > 0 {
		ptn = (*C.HPDF_UINT16)(&dashPtn[0])
	}
	C.HPDF_Page_SetDash(page.page, ptn, C.HPDF_UINT(len(dashPtn)), C.HPDF_UINT(phase))
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

func (page *Page) SetLineWidth(lineWidth float32) error {
	C.HPDF_Page_SetLineWidth(page.page, C.HPDF_REAL(lineWidth))
  
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

func (page *Page) GetHorizontalScalling() float32 {
	return float32(C.HPDF_Page_GetHorizontalScalling(page.page))
}

func (page *Page) SetHorizontalScalling(scale float32) error {
	C.HPDF_Page_SetHorizontalScalling(page.page, C.HPDF_REAL(scale))

  return page.pdf.GetLastError()
}

func (page *Page) Stroke() error {
	C.HPDF_Page_Stroke(page.page)

	return page.pdf.GetLastError()
}
