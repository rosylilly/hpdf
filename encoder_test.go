package hpdf

import (
	. "testing"
)

func TestGetEncoder(t *T) {
	pdf, _ := New()

	encoder, err := pdf.GetEncoder("StandardEncoding")

	if encoder == nil || err != nil {
		t.Fatal(err)
	}

	_, err = pdf.GetEncoder("a invalid encoding name")

	if err == nil {
		t.Fatal("Not error occured")
	}
}

func TestGetCurrentEncoder(t *T) {
	pdf, _ := New()

	err := pdf.SetCurrentEncoder("StandardEncoding")

	if err != nil {
		t.Fatal(err)
	}

	encoder := pdf.GetCurrentEncoder()

	if encoder == nil {
		t.Fatal(pdf.GetLastError())
	}
}
