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
