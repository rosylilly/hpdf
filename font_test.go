package hpdf

import (
	. "testing"
)

func TestGetFont(t *T) {
	pdf, _ := New()

	font, err := pdf.GetFont("Courier")

	if font == nil || err != nil {
		t.Fatal(err)
	}
}

func TestLoadType1FontFromFile(t *T) {
	t.Skip("I don't have afm files")
}

func TestLoadTTFontFromFile(t *T) {
	t.Skip("I dont't have ttf files")
}

func TestLoadTTFontFromFile2(t *T) {
	t.Skip("I dont't have ttc files")
}

func TestUseJPFonts(t *T) {
	pdf, _ := New()

	err := pdf.UseJPFonts()

	if err != nil {
		t.Fatal(err)
	}
}

func TestUseKRFonts(t *T) {
	pdf, _ := New()

	err := pdf.UseKRFonts()

	if err != nil {
		t.Fatal(err)
	}
}

func TestUseCNSFonts(t *T) {
	pdf, _ := New()

	err := pdf.UseCNSFonts()

	if err != nil {
		t.Fatal(err)
	}
}

func TestUseCNTFonts(t *T) {
	pdf, _ := New()

	err := pdf.UseCNTFonts()

	if err != nil {
		t.Fatal(err)
	}
}
