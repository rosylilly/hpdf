package hpdf

import (
	. "testing"
)

func CreatePage() (*Page, *PDF, error) {
	pdf, _ := New()
	page, err := pdf.AddPage()

	return page, pdf, err
}

func SetFont(page *Page, pdf *PDF, fontName string) error {
	font, _ := pdf.GetFont(fontName)
	return page.SetFontAndSize(font, 6)
}

func TestBeginAndEndText(t *T) {
	page, _, err := CreatePage()
	if err != nil {
		t.Fatal(err)
	}

	err = page.BeginText()

	if err != nil {
		t.Fatal(err)
	}

	err = page.EndText()

	if err != nil {
		t.Fatal(err)
	}
}

func TestShowText(t *T) {
	page, pdf, err := CreatePage()

	if err != nil {
		t.Fatal(err)
	}

	err = SetFont(page, pdf, "Courier")
	if err != nil {
		t.Fatal(err)
	}

	page.BeginText()

	err = page.ShowText("Some Text")

	if err != nil {
		t.Fatal(err)
	}

	page.EndText()

}

func TestTextOut(t *T) {
	page, pdf, err := CreatePage()
	if err != nil {
		t.Fatal(err)
	}

	err = SetFont(page, pdf, "Courier")
	if err != nil {
		t.Fatal(err)
	}

	page.BeginText()

	err = page.TextOut(0, 0, "Some Text")

	if err != nil {
		t.Fatal(err)
	}

	page.EndText()

}
