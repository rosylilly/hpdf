package hpdf

import (
	. "testing"
)

func TestCreateDestionation(t *T) {
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

func TestSetXYZ(t *T) {
	pdf, _ := New()
	page, _ := pdf.AddPage()

	var d *Destination

	d, _ = page.CreateDestination()
	err := d.SetXYZ(0, 0, 1.00)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSetFit(t *T) {
	pdf, _ := New()
	page, _ := pdf.AddPage()

	var d *Destination

	d, _ = page.CreateDestination()
	err := d.SetFit()

	if err != nil {
		t.Fatal(err)
	}
}

func TestSetFitH(t *T) {
	pdf, _ := New()
	page, _ := pdf.AddPage()

	var d *Destination

	d, _ = page.CreateDestination()
	err := d.SetFitH(0)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSetFitV(t *T) {
	pdf, _ := New()
	page, _ := pdf.AddPage()

	var d *Destination

	d, _ = page.CreateDestination()
	err := d.SetFitV(0)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSetFitR(t *T) {
	pdf, _ := New()
	page, _ := pdf.AddPage()

	var d *Destination

	d, _ = page.CreateDestination()
	err := d.SetFitR(0, 0, 0, 0)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSetFitB(t *T) {
	pdf, _ := New()
	page, _ := pdf.AddPage()

	var d *Destination

	d, _ = page.CreateDestination()
	err := d.SetFitB()

	if err != nil {
		t.Fatal(err)
	}
}

func TestSetFitBH(t *T) {
	pdf, _ := New()
	page, _ := pdf.AddPage()

	var d *Destination

	d, _ = page.CreateDestination()
	err := d.SetFitBH(0)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSetFitBV(t *T) {
	pdf, _ := New()
	page, _ := pdf.AddPage()

	var d *Destination

	d, _ = page.CreateDestination()
	err := d.SetFitBV(0)

	if err != nil {
		t.Fatal(err)
	}
}
