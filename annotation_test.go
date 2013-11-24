package hpdf

import (
	. "testing"
)

func TestCreateTextAnnot(t *T) {
	pdf, _ := New()
	page, _ := pdf.AddPage()

	rect := NewRect(0, 0, 0, 0)
	annotation, err := page.CreateTextAnnot(rect, "test", nil)

	if annotation == nil || err != nil {
		t.Fatal(err)
	}
}

func TestCreateLinkAnnot(t *T) {
	pdf, _ := New()
	page, _ := pdf.AddPage()

	rect := NewRect(0, 0, 0, 0)
	destination, _ := page.CreateDestination()
	annotation, err := page.CreateLinkAnnot(rect, destination)

	if annotation == nil || err != nil {
		t.Fatal(err)
	}
}

func TestCreateURILinkAnnot(t *T) {
	pdf, _ := New()
	page, _ := pdf.AddPage()

	rect := NewRect(0, 0, 0, 0)
	annotation, err := page.CreateURILinkAnnot(rect, "http://www.libharu.org/")

	if annotation == nil || err != nil {
		t.Fatal(err)
	}
}

func TestAnnotSetBorderStyle(t *T) {
	pdf, _ := New()
	page, _ := pdf.AddPage()

	rect := NewRect(0, 0, 0, 0)
	annotation, err := page.CreateURILinkAnnot(rect, "http://www.libharu.org/")

	if annotation == nil || err != nil {
		t.Fatal(err)
	}

	err = annotation.SetBorderStyle(BS_UNDERLINED, 100, 1, 1, 1)

	if err != nil {
		t.Fatal(err)
	}
}

func TestTextAnnotSetIcon(t *T) {
	pdf, _ := New()
	page, _ := pdf.AddPage()

	rect := NewRect(0, 0, 0, 0)
	annotation, err := page.CreateTextAnnot(rect, "test", nil)

	if annotation == nil || err != nil {
		t.Fatal(err)
	}

	err = annotation.SetIcon(ANNOT_ICON_INSERT)

	if err != nil {
		t.Fatal(err)
	}
}

func TestTextAnnotSetOpened(t *T) {
	pdf, _ := New()
	page, _ := pdf.AddPage()

	rect := NewRect(0, 0, 0, 0)
	annotation, err := page.CreateTextAnnot(rect, "test", nil)

	if annotation == nil || err != nil {
		t.Fatal(err)
	}

	err = annotation.SetOpened(true)

	if err != nil {
		t.Fatal(err)
	}
}

func TestLinkAnnotSetHighlightMode(t *T) {
	pdf, _ := New()
	page, _ := pdf.AddPage()

	rect := NewRect(0, 0, 0, 0)
	destination, _ := page.CreateDestination()
	annotation, err := page.CreateLinkAnnot(rect, destination)

	if annotation == nil || err != nil {
		t.Fatal(err)
	}

	err = annotation.SetHighlightMode(ANNOT_DOWN_APPEARANCE)

	if err != nil {
		t.Fatal(err)
	}
}

func TestLinkAnnotSetBorderStyle(t *T) {
	pdf, _ := New()
	page, _ := pdf.AddPage()

	rect := NewRect(0, 0, 0, 0)
	destination, _ := page.CreateDestination()
	annotation, err := page.CreateLinkAnnot(rect, destination)

	if annotation == nil || err != nil {
		t.Fatal(err)
	}

	err = annotation.SetBorderStyle(100, 1, 1)

	if err != nil {
		t.Fatal(err)
	}
}
