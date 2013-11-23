package hpdf

import (
	. "testing"
)

func TestNewPDF(t *T) {
	_, err := New()

	if err != nil {
		t.Fatal(err)
	}
}

func TestPDFFree(t *T) {
	pdf, err := New()

	if err != nil {
		t.Fatal(err)
	}

	pdf.Free()
}
