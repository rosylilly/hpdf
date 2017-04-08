package hpdf

import (
	. "testing"
)

func TestAddPage(t *T) {
	pdf, _ := New()

	page, err := pdf.AddPage()
	if page == nil || err != nil {
		t.Fatal(err)
	}
}

func TestInsertPage(t *T) {
	pdf, _ := New()

	page, err := pdf.AddPage()
	if err != nil {
		t.Fatal(err)
	}

	insertedPage, err := pdf.InsertPage(page)

	if insertedPage == nil || err != nil {
		t.Fatal(err)
	}
}

func TestPageSetWidth(t *T) {
	pdf, _ := New()

	page, err := pdf.AddPage()
	if err != nil {
		t.Fatal(err)
	}

	width := float32(200)

	err = page.SetWidth(width)

	if err != nil || page.GetWidth() != width {
		t.Fatal(err)
	}
}

func TestPageSetHeight(t *T) {
	pdf, _ := New()

	page, err := pdf.AddPage()
	if err != nil {
		t.Fatal(err)
	}

	height := float32(200)

	err = page.SetHeight(height)

	if err != nil || page.GetHeight() != height {
		t.Fatal(err)
	}
}

func TestPageSetSize(t *T) {
	pdf, _ := New()

	page, err := pdf.AddPage()
	if err != nil {
		t.Fatal(err)
	}

	err = page.SetSize(PAGE_SIZE_COMM10, PAGE_LANDSCAPE)

	if err != nil {
		t.Fatal(err)
	}
}

func TestPageSetRotate(t *T) {
	pdf, _ := New()

	page, err := pdf.AddPage()
	if err != nil {
		t.Fatal(err)
	}

	err = page.SetRotate(90)

	if err != nil {
		t.Fatal(err)
	}
}

func TestPageTextWidth(t *T) {
	pdf, _ := New()
	font, _ := pdf.GetFont("Courier")

	page, err := pdf.AddPage()
	if err != nil {
		t.Fatal(err)
	}

	page.SetFontAndSize(font, 12)

	width := page.TextWidth("text")

	if width == 0 {
		t.Fatal(pdf.GetLastError())
	}
}

func TestPageMesureText(t *T) {
	pdf, _ := New()
	font, _ := pdf.GetFont("Courier")

	page, err := pdf.AddPage()
	if err != nil {
		t.Fatal(err)
	}

	page.SetFontAndSize(font, 12)

	width, err := page.MeasureText("text", 100, false)

	if err != nil {
		t.Fatal(err)
	}

	if width != 28.8 {
		t.Fatalf("Broken width: %f", width)
	}

	width, err = page.MeasureText("text text", 30, true)

	if err != nil {
		t.Fatal(err)
	}

	if width <= 0 {
		t.Fatalf("Broken width: %f", width)
	}
}

func TestPageGetGMode(t *T) {
	pdf, _ := New()

	page, err := pdf.AddPage()
	if err != nil {
		t.Fatal(err)
	}

	if gMode := page.GetGMode(); gMode != GMODE_PAGE_DESCRIPTION {
		t.Fatalf("Broken width: %v", gMode)
	}
}

func TestPageGetCurrentPos(t *T) {
	pdf, _ := New()

	page, err := pdf.AddPage()
	if err != nil {
		t.Fatal(err)
	}

	point := page.GetCurrentPos()

	if point == nil {
		t.Fatal("broken GetCurrentPos")
	}
}

func TestPageGetCurrentTextPos(t *T) {
	pdf, _ := New()

	page, err := pdf.AddPage()
	if err != nil {
		t.Fatal(err)
	}

	point := page.GetCurrentTextPos()

	if point == nil {
		t.Fatal("broken GetCurrentTextPos")
	}
}

func TestPageGetCurrentFont(t *T) {
	pdf, _ := New()
	font, _ := pdf.GetFont("Courier")

	page, err := pdf.AddPage()
	if err != nil {
		t.Fatal(err)
	}

	if actual := page.GetCurrentFont(); actual != nil {
		t.Fatalf("Got: %+v", actual)
	}

	page.SetFontAndSize(font, 12)

	actual := page.GetCurrentFont()
	if actual == nil {
		t.Fatal("Got nil")
	}
}

func TestPageGetCurrentFontSize(t *T) {
	pdf, _ := New()
	font, _ := pdf.GetFont("Courier")

	page, err := pdf.AddPage()
	if err != nil {
		t.Fatal(err)
	}

	if actual := page.GetCurrentFontSize(); actual != 0 {
		t.Fatalf("Got: %+v", actual)
	}

	page.SetFontAndSize(font, 12)

	actual := page.GetCurrentFontSize()
	if actual == 0 {
		t.Fatal("Got zero")
	}
}

func TestPageGetTransMatrix(t *T) {
	pdf, _ := New()

	page, err := pdf.AddPage()
	if err != nil {
		t.Fatal(err)
	}

	transMatrix := page.GetTransMatrix()
	if transMatrix == nil {
		t.Fatal(pdf.GetLastError())
	}
}

func TestPageGetLineWidth(t *T) {
	pdf, _ := New()

	page, err := pdf.AddPage()
	if err != nil {
		t.Fatal(err)
	}

	lw := page.GetLineWidth()

	if lw == 0 {
		t.Fatal(pdf.GetLastError())
	}
}

func TestPageGetLineCap(t *T) {
	pdf, _ := New()

	page, err := pdf.AddPage()
	if err != nil {
		t.Fatal(err)
	}

	lc := page.GetLineCap()

	if lc != LINE_CAP_BUTT_END {
		t.Fatal(lc)
	}
}

func TestPageSetLineCap(t *T) {
	pdf, _ := New()

	page, err := pdf.AddPage()
	if err != nil {
		t.Fatal(err)
	}

	err = page.SetLineCap(LINE_CAP_ROUND_END)

	if err != nil {
		t.Fatal(err)
	}

	lc := page.GetLineCap()
	if lc != LINE_CAP_ROUND_END {
		t.Fatal(lc)
	}
}

func TestPageGetLineJoin(t *T) {
	pdf, _ := New()

	page, err := pdf.AddPage()
	if err != nil {
		t.Fatal(err)
	}

	lj := page.GetLineJoin()

	if lj != LINE_JOIN_MITER_JOIN {
		t.Fatal(lj)
	}
}

func TestPageGetMiterLimit(t *T) {
	pdf, _ := New()

	page, err := pdf.AddPage()
	if err != nil {
		t.Fatal(err)
	}

	ml := page.GetMiterLimit()

	if err := pdf.GetLastError(); err != nil {
		t.Log(ml)
		t.Fatal(err)
	}
}

func TestPageGetDash(t *T) {
	pdf, _ := New()

	page, err := pdf.AddPage()
	if err != nil {
		t.Fatal(err)
	}

	dm := page.GetDash()

	err = pdf.GetLastError()

	if err != nil || dm == nil {
		t.Fatal(err, dm)
	}
	t.Logf("%+v", dm)
}
