package main

import (
	hpdf "github.com/rosylilly/hpdf"
)

func main() {
	pdf, _ := hpdf.New()
	page, _ := pdf.AddPage()
	image, _ := pdf.LoadPngImageFromFile("../testdata/png.png")
	page.SetWidth(470)
	page.SetHeight(470)
	page.DrawImage(image, 0, 0, 470, 470)
	pdf.SaveToFile("example.pdf")
}
