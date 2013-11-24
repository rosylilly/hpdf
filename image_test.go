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
