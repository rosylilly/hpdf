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

func TestPDFSaveToFile(t *T) {
	pdf, err := New()

	if err != nil {
		t.Fatal(err)
	}

	err = pdf.SaveToFile("tmp/save_to_file.pdf")

	if err != nil {
		t.Fatal(err)
	}
}
