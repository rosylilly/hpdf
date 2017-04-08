package hpdf

import (
	. "testing"
)

func TestPageCurveTo(t *T) {
	pdf, _ := New()
	page, _ := pdf.AddPage()

	page.MoveTo(0.5, 0.6)

	err := page.CurveTo(1.1, 2.2, 3.3, 4.4, 5.5, 6.6)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPageDrawImage(t *T) {
	pdf, _ := New()
	page, _ := pdf.AddPage()
	image, _ := pdf.LoadPngImageFromFile("testdata/png.png")

	err := page.DrawImage(image, 0, 0, 470, 470)

	if err != nil {
		t.Fatal(err)
	}
}

func TestPageFill(t *T) {
	pdf, _ := New()
	page, _ := pdf.AddPage()

	page.Rectangle(1.5, 2.5, 3.5, 4.5)
	err := page.Fill()
	if err != nil {
		t.Fatal(err)
	}
}

func TestPageFillStroke(t *T) {
	pdf, _ := New()
	page, _ := pdf.AddPage()

	page.Rectangle(1.5, 2.5, 3.5, 4.5)
	err := page.FillStroke()
	if err != nil {
		t.Fatal(err)
	}
}

func TestPageLineTo(t *T) {
	pdf, _ := New()
	page, _ := pdf.AddPage()

	page.MoveTo(1.5, 2.5)
	err := page.LineTo(3.5, 4.5)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPageMoveTo(t *T) {
	pdf, _ := New()
	page, _ := pdf.AddPage()

	err := page.MoveTo(0.5, 0.6)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPageRectangle(t *T) {
	pdf, _ := New()
	page, _ := pdf.AddPage()

	err := page.Rectangle(1.5, 2.5, 3.5, 4.5)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPageSetDash(t *T) {
	pdf, _ := New()
	page, _ := pdf.AddPage()

	err := page.SetDash([]uint16{2, 3}, 1)
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

func TestPageSetLineWidth(t *T) {
	pdf, _ := New()
	page, _ := pdf.AddPage()

	err := page.SetLineWidth(2.5)
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

func TestPageSetRGBStroke(t *T) {
	pdf, _ := New()
	page, _ := pdf.AddPage()

	err := page.SetRGBStroke(0.50, 0.32, 0.75)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPageSetCMYKFill(t *T) {
	pdf, _ := New()
	page, _ := pdf.AddPage()

	err := page.SetCMYKFill(0.50, 0.32, 0.75, 0.25)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPageSetCMYKStroke(t *T) {
	pdf, _ := New()
	page, _ := pdf.AddPage()

	err := page.SetCMYKStroke(0.50, 0.32, 0.75, 0.25)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPageStroke(t *T) {
	pdf, _ := New()
	page, _ := pdf.AddPage()

	page.MoveTo(1.5, 2.5)
	page.LineTo(3.5, 4.5)
	err := page.Stroke()
	if err != nil {
		t.Fatal(err)
	}
}
