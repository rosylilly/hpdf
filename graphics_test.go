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

	c := page.GetRGBFill()

	if c.R != 0.50 || c.G != 0.32 || c.B != 0.75 {
		t.Fatalf("Invalid color: %+v", c)
	}
}

func TestPageSetRGBStroke(t *T) {
	pdf, _ := New()
	page, _ := pdf.AddPage()

	err := page.SetRGBStroke(0.50, 0.32, 0.75)
	if err != nil {
		t.Fatal(err)
	}

	c := page.GetRGBStroke()

	if c.R != 0.50 || c.G != 0.32 || c.B != 0.75 {
		t.Fatalf("Invalid color: %+v", c)
	}
}

func TestPageSetCMYKFill(t *T) {
	pdf, _ := New()
	page, _ := pdf.AddPage()

	err := page.SetCMYKFill(0.50, 0.32, 0.75, 0.25)
	if err != nil {
		t.Fatal(err)
	}

	c := page.GetCMYKFill()

	if c.C != 0.50 || c.M != 0.32 || c.Y != 0.75 || c.K != 0.25 {
		t.Fatalf("Invalid color: %+v", c)
	}
}

func TestPageSetCMYKStroke(t *T) {
	pdf, _ := New()
	page, _ := pdf.AddPage()

	err := page.SetCMYKStroke(0.50, 0.32, 0.75, 0.25)
	if err != nil {
		t.Fatal(err)
	}

	c := page.GetCMYKStroke()

	if c.C != 0.50 || c.M != 0.32 || c.Y != 0.75 || c.K != 0.25 {
		t.Fatalf("Invalid color: %+v", c)
	}
}
