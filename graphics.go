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
