package hpdf

/*
#cgo LDFLAGS: -lhpdf -lpng -lz
#include "hpdf.h"
*/
import "C"
import (
	"unsafe"
)

func (pdf *PDF) LoadPngImageFromFile(filename string) (*Image, error) {
	cfilename := C.CString(filename)
	cimage := C.HPDF_LoadPngImageFromFile(pdf.doc, cfilename)
	C.free(unsafe.Pointer(cfilename))

	if cimage != nil {
		return newImage(cimage, pdf), nil
	} else {
		return nil, pdf.GetLastError()
	}
}

func (pdf *PDF) LoadPngImageFromFile2(filename string) (*Image, error) {
	cfilename := C.CString(filename)
	cimage := C.HPDF_LoadPngImageFromFile2(pdf.doc, cfilename)
	C.free(unsafe.Pointer(cfilename))

	if cimage != nil {
		return newImage(cimage, pdf), nil
	} else {
		return nil, pdf.GetLastError()
	}
}

func (pdf *PDF) LoadRawImageFromFile(
	filename string, width, height uint32, colorSpace ColorSpace,
) (*Image, error) {
	cfilename := C.CString(filename)
	cimage := C.HPDF_LoadRawImageFromFile(
		pdf.doc, cfilename,
		C.HPDF_UINT(width), C.HPDF_UINT(height),
		C.HPDF_ColorSpace(colorSpace),
	)
	C.free(unsafe.Pointer(cfilename))

	if cimage != nil {
		return newImage(cimage, pdf), nil
	} else {
		return nil, pdf.GetLastError()
	}
}

func (pdf *PDF) LoadJpegImageFromFile(filename string) (*Image, error) {
	cfilename := C.CString(filename)
	cimage := C.HPDF_LoadJpegImageFromFile(pdf.doc, cfilename)
	C.free(unsafe.Pointer(cfilename))

	if cimage != nil {
		return newImage(cimage, pdf), nil
	} else {
		return nil, pdf.GetLastError()
	}
}

type Image struct {
	image C.HPDF_Image
	pdf *PDF
}

func newImage(image C.HPDF_Image, pdf *PDF) *Image {
	return &Image{image, pdf}
}
