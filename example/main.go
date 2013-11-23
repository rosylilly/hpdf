package main

import (
	hpdf "github.com/rosylilly/hpdf"
)

func main() {
	pdf, _ := hpdf.New()
	pdf.AddPage()
	pdf.AddPage()
	pdf.AddPage()
	pdf.SaveToFile("example.pdf")
}
