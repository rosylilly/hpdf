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

func TestPageCreateDestionation(t *T) {
	pdf, _ := New()

	page, err := pdf.AddPage()
	if err != nil {
		t.Fatal(err)
	}

	destination, err := page.CreateDestination()

	if destination == nil || err != nil {
		t.Fatal(err)
	}
}
