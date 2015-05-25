package hpdf

import (
	. "testing"
)

func TestPageDrawImage(t *T) {
	pdf, _ := New()
	page, _ := pdf.AddPage()
	image, _ := pdf.LoadPngImageFromFile("testdata/png.png")

	err := page.DrawImage(image, 0, 0, 470, 470)

	if err != nil {
		t.Fatal(err)
	}
}

func TestPageSetFontAndSize(t *T) {
	pdf, _ := New()
	font, _ := pdf.GetFont("Courier")
	page, _ := pdf.AddPage()

	err := page.SetFontAndSize(font, 12)

	if err != nil {
		t.Fatal(err)
	}
}

func TestPageSetRGBFill(t *T) {
	pdf, _ := New()
	page, _ := pdf.AddPage()

	err := page.SetRGBFill(0.50, 0.32, 0.75)
	if err != nil {
		t.Fatal(err)
	}
}
