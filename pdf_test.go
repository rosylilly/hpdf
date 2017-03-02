package hpdf

import (
	"os"
	. "testing"
	"time"
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

	f, err := os.Create("tmp/save_to_stream.pdf")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	err = pdf.SaveToStream(f)

	if err != nil {
		t.Fatal(err)
	}
}

func TestPDFHasDoc(t *T) {
	pdf, err := New()

	if err != nil {
		t.Fatal(err)
	}

	if !pdf.HasDoc() {
		t.Fatal(pdf.GetLastError())
	}

	pdf.FreeDocAll()

	if pdf.HasDoc() {
		t.Fatal(pdf.GetLastError())
	}

	pdf.NewDoc()

	if !pdf.HasDoc() {
		t.Fatal(pdf.GetLastError())
	}
}

func TestPDFSetPageConfiguration(t *T) {
	pdf, err := New()

	if err != nil {
		t.Fatal(err)
	}

	err = pdf.SetPagesConfiguration(100)

	if err != nil {
		t.Fatal(err)
	}
}

func TestPDFSetPageLayout(t *T) {
	pdf, err := New()

	if err != nil {
		t.Fatal(err)
	}

	err = pdf.SetPageLayout(PAGE_LAYOUT_TWO_COLUMN_RIGHT)

	if err != nil {
		t.Fatal(err)
	}

	if pdf.GetPageLayout() != PAGE_LAYOUT_TWO_COLUMN_RIGHT {
		t.Fatalf("Missmatch page layout: %d", pdf.GetPageLayout())
	}
}

func TestPDFSetPageMode(t *T) {
	pdf, err := New()

	if err != nil {
		t.Fatal(err)
	}

	err = pdf.SetPageMode(PAGE_MODE_FULL_SCREEN)

	if err != nil {
		t.Fatal(err)
	}

	if pdf.GetPageMode() != PAGE_MODE_FULL_SCREEN {
		t.Fatalf("Missmatch page layout: %d", pdf.GetPageMode())
	}
}

func TestPDFSetOpenAction(t *T) {
	pdf, err := New()

	if err != nil {
		t.Fatal(err)
	}

	page, _ := pdf.AddPage()
	destination, _ := page.CreateDestination()

	err = pdf.SetOpenAction(destination)

	if err != nil {
		t.Fatal(err)
	}
}

func TestPDFAddPageLabel(t *T) {
	pdf, err := New()

	if err != nil {
		t.Fatal(err)
	}

	err = pdf.AddPageLabel(0, PAGE_NUM_STYLE_LOWER_LETTERS, 0, "test-")

	if err != nil {
		t.Fatal(err)
	}
}

func TestPDFSetInfoAttr(t *T) {
	pdf, err := New()

	if err != nil {
		t.Fatal(err)
	}

	err = pdf.SetInfoAttr(INFO_TITLE, "HPDF Test")

	if err != nil {
		t.Fatal(err)
	}

	title, err := pdf.GetInfoAttr(INFO_TITLE)

	if title != "HPDF Test" || err != nil {
		t.Fatal(err)
	}

	date := time.Now().Add(10 * time.Hour)
	err = pdf.SetInfoDateAttr(INFO_CREATION_DATE, date)

	if err != nil {
		t.Fatal(err)
	}
}

func TestPDFSetPassword(t *T) {
	pdf, err := New()

	if err != nil {
		t.Fatal(err)
	}

	err = pdf.SetPassword("owner pass", "user pass")

	if err != nil {
		t.Fatal(err)
	}
}

func TestPDFSetPermission(t *T) {
	pdf, err := New()

	if err != nil {
		t.Fatal(err)
	}

	err = pdf.SetPassword("owner pass", "user pass")

	if err != nil {
		t.Fatal(err)
	}

	err = pdf.SetPermission(ENABLE_READ | ENABLE_COPY)

	if err != nil {
		t.Fatal(err)
	}
}

func TestPDFSetViewerPreference(t *T) {
	pdf, err := New()

	if err != nil {
		t.Fatal(err)
	}

	pref := pdf.GetViewerPreference()

	if pref != 0 {
		t.Fatalf("Viewer Preferece: %d", pref)
	}

	pdf.SetViewerPreference(HIDE_TOOLBAR | HIDE_MENUBAR)

	pref = pdf.GetViewerPreference()

	if pref != 3 {
		t.Fatalf("Viewer Preferece: %d", pref)
	}
}

func TestPDFSetEncryptMode(t *T) {
	pdf, err := New()

	if err != nil {
		t.Fatal(err)
	}

	err = pdf.SetPassword("owner pass", "user pass")

	if err != nil {
		t.Fatal(err)
	}

	err = pdf.SetEncryptMode(ENCRYPT_R3, 10)

	if err != nil {
		t.Fatal(err)
	}
}

func TestPDFSetCompressionMode(t *T) {
	pdf, err := New()

	if err != nil {
		t.Fatal(err)
	}

	err = pdf.SetCompressionMode(COMP_ALL)

	if err != nil {
		t.Fatal(err)
	}
}
