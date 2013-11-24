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

func TestFontGetFontName(t *T) {
	pdf, _ := New()

	font, err := pdf.GetFont("Courier")

	if font == nil || err != nil {
		t.Fatal(err)
	}

	fontName := font.GetFontName()

	if fontName != "Courier" {
		t.Fatalf("Font name missmatch: %s", fontName)
	}
}

func TestFontGetEncodingName(t *T) {
	pdf, _ := New()

	font, err := pdf.GetFont("Courier")

	if font == nil || err != nil {
		t.Fatal(err)
	}

	encodingName := font.GetEncodingName()

	if encodingName != "StandardEncoding" {
		t.Fatalf("Encoding name missmatch: %s", encodingName)
	}
}

func TestFontGetUnicodeWidth(t *T) {
	pdf, _ := New()

	font, err := pdf.GetFont("Courier")

	if font == nil || err != nil {
		t.Fatal(err)
	}

	width := font.GetUnicodeWidth('a')

	if width != 600 {
		t.Fatal("Missmatch unicode width:", width)
	}
}

func TestFontGetBBox(t *T) {
	pdf, _ := New()

	font, err := pdf.GetFont("Courier")

	if font == nil || err != nil {
		t.Fatal(err)
	}

	box := font.GetBBox()

	if box.Left == 0 && box.Bottom == 0 && box.Right == 0 && box.Top == 0 {
		t.Fatalf("%+v", box)
	}
}

func TestFontGetAscent(t *T) {
	pdf, _ := New()

	font, err := pdf.GetFont("Courier")

	if font == nil || err != nil {
		t.Fatal(err)
	}

	ascent := font.GetAscent()

	if ascent == 0 {
		t.Fatal("Invalid ascent:", ascent)
	}
}

func TestFontGetDescent(t *T) {
	pdf, _ := New()

	font, err := pdf.GetFont("Courier")

	if font == nil || err != nil {
		t.Fatal(err)
	}

	descent := font.GetDescent()

	if descent == 0 {
		t.Fatal("Invalid descent:", descent)
	}
}

func TestFontGetXHeight(t *T) {
	pdf, _ := New()

	font, err := pdf.GetFont("Courier")

	if font == nil || err != nil {
		t.Fatal(err)
	}

	xHeight := font.GetXHeight()

	if xHeight == 0 {
		t.Fatal("Invalid xHeight:", xHeight)
	}
}

func TestFontGetCapHeight(t *T) {
	pdf, _ := New()

	font, err := pdf.GetFont("Courier")

	if font == nil || err != nil {
		t.Fatal(err)
	}

	capHeight := font.GetCapHeight()

	if capHeight == 0 {
		t.Fatal("Invalid capHeight:", capHeight)
	}
}
