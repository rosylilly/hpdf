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
	pdf, _ := New()
	pdf.Free()
}

func TestPDFNewDoc(t *T) {
	pdf, _ := New()

	err := pdf.NewDoc()

	if err != nil {
		t.Fatal(err)
	}
}

func TestPDFFreeDoc(t *T) {
	pdf, _ := New()
	pdf.NewDoc()
	pdf.FreeDoc()

	if err := pdf.GetLastError(); err != nil {
		t.Fatal(err)
	}
}

func TestPDFFreeDocAll(t *T) {
	pdf, _ := New()
	pdf.NewDoc()
	pdf.FreeDocAll()

	if err := pdf.GetLastError(); err != nil {
		t.Fatal(err)
	}
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

func TestPDFSaveToStream(t *T) {
	pdf, err := New()

	if err != nil {
		t.Fatal(err)
	}

	err = pdf.SaveToStream()

	if err != nil {
		t.Fatal(err)
	}
}

func TestPDFGetStreamSize(t *T) {
	pdf, err := New()

	if err != nil {
		t.Fatal(err)
	}

	if pdf.GetStreamSize() != 0 {
		t.Fatal(pdf.GetLastError())
	}
}

func TestPDFResetStream(t *T) {
	pdf, err := New()

	if err != nil {
		t.Fatal(err)
	}

	err = pdf.ResetStream()
	hpdfError := err.(*Error)

	if hpdfError.Code != ERR_INVALID_OPERATION {
		t.Fatal(err)
	}
}

func TestPDFHasDoc(t *T) {
	pdf, err := New()

	if err != nil {
		t.Fatal(err)
	}

	if exist, err := pdf.HasDoc(); !exist {
		t.Fatal(err)
	}

	pdf.FreeDocAll()

	if exist, err := pdf.HasDoc(); exist {
		t.Fatal(err)
	}

	pdf.NewDoc()

	if exist, err := pdf.HasDoc(); !exist {
		t.Fatal(err)
	}
}
