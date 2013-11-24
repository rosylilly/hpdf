package hpdf

import (
	. "testing"
)

func TestLoadPngImageFromFile(t *T) {
	pdf, _ := New()

	png, err := pdf.LoadPngImageFromFile("testdata/png.png")

	if png == nil || err != nil {
		t.Fatal(err)
	}
}

func TestLoadPngImageFromFile2(t *T) {
	pdf, _ := New()

	png, err := pdf.LoadPngImageFromFile2("testdata/png.png")

	if png == nil || err != nil {
		t.Fatal(err)
	}
}

func TestLoadRawImageFromFile(t *T) {
	pdf, _ := New()

	raw, err := pdf.LoadRawImageFromFile("testdata/raw.raw", 470, 470, CS_DEVICE_RGB)

	if raw == nil || err != nil {
		t.Fatal(err)
	}
}

func TestLoadJpegImageFromFile(t *T) {
	pdf, _ := New()

	jpeg, err := pdf.LoadJpegImageFromFile("testdata/jpg.jpg")

	if jpeg == nil || err != nil {
		t.Fatal(err)
	}
}

func TestImageGetSize(t *T) {
	pdf, _ := New()

	png, err := pdf.LoadPngImageFromFile("testdata/png.png")

	if png == nil || err != nil {
		t.Fatal(err)
	}

	point := png.GetSize()

	if point.X != 470 || point.Y != 470 {
		t.Fatal("Missmatch size")
	}
}

func TestImageGetWidth(t *T) {
	pdf, _ := New()

	png, err := pdf.LoadPngImageFromFile("testdata/png.png")

	if png == nil || err != nil {
		t.Fatal(err)
	}

	width := png.GetWidth()

	if width != 470 {
		t.Fatal("Missmatch width:", width)
	}
}

func TestImageGetHeight(t *T) {
	pdf, _ := New()

	png, err := pdf.LoadPngImageFromFile("testdata/png.png")

	if png == nil || err != nil {
		t.Fatal(err)
	}

	height := png.GetHeight()

	if height != 470 {
		t.Fatal("Missmatch height:", height)
	}
}

func TestImageGetBitsPerComponent(t *T) {
	pdf, _ := New()

	png, err := pdf.LoadPngImageFromFile("testdata/png.png")

	if png == nil || err != nil {
		t.Fatal(err)
	}

	bitsPerComponent := png.GetBitsPerComponent()

	if bitsPerComponent != 8 {
		t.Fatal("Missmatch bitsPerComponent:", bitsPerComponent)
	}
}

func TestImageGetColorSpace(t *T) {
	pdf, _ := New()

	png, err := pdf.LoadPngImageFromFile("testdata/png.png")

	if png == nil || err != nil {
		t.Fatal(err)
	}

	colorSpace := png.GetColorSpace()

	if colorSpace != CS_DEVICE_RGB {
		t.Fatal("Missmatch color space:", colorSpace)
	}
}

func TestImageSetColorMask(t *T) {
	pdf, _ := New()

	png, err := pdf.LoadPngImageFromFile("testdata/png.png")

	if png == nil || err != nil {
		t.Fatal(err)
	}

	mask, err := pdf.LoadPngImageFromFile("testdata/mask.png")

	if mask == nil || err != nil {
		t.Fatal(err)
	}

	t.Log(png.GetBitsPerComponent())
	t.Log(mask.GetBitsPerComponent())

	err = png.SetMaskImage(mask)

	if err != nil {
		t.Fatal(err)
	}
}
