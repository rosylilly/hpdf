package hpdf

import (
	. "testing"
)

func TestCreateOutline(t *T) {
	pdf, _ := New()

	outline, err := pdf.CreateOutline(nil, "root", nil)

	if outline == nil || err != nil {
		t.Fatal(err)
	}
}

func TestOutlineSetOpened(t *T) {
	pdf, _ := New()

	outline, err := pdf.CreateOutline(nil, "root", nil)

	if outline == nil || err != nil {
		t.Fatal(err)
	}

	err = outline.SetOpened(true)

	if err != nil {
		t.Fatal(err)
	}
}

func TestOutlineSetDestination(t *T) {
	pdf, _ := New()
	page, _ := pdf.AddPage()
	destination, _ := page.CreateDestination()

	outline, err := pdf.CreateOutline(nil, "root", nil)

	if outline == nil || err != nil {
		t.Fatal(err)
	}

	outline.SetDestination(destination)

	if err != nil {
		t.Fatal(err)
	}
}
