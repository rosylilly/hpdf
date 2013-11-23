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
