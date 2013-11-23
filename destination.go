package hpdf

/*
#cgo LDFLAGS: -lhpdf -lpng -lz
#include "hpdf.h"
*/
import "C"

type Destination struct {
	destination C.HPDF_Destination
	page        *Page
}

func (destination *Destination) SetXYZ(left, top, zoom float32) error {
	C.HPDF_Destination_SetXYZ(
		destination.destination,
		C.HPDF_REAL(left), C.HPDF_REAL(top), C.HPDF_REAL(zoom),
	)

	return destination.page.pdf.GetLastError()
}

func (destination *Destination) SetFit() error {
	C.HPDF_Destination_SetFit(destination.destination)

	return destination.page.pdf.GetLastError()
}

func (destination *Destination) SetFitH(top float32) error {
	C.HPDF_Destination_SetFitH(destination.destination, C.HPDF_REAL(top))

	return destination.page.pdf.GetLastError()
}

func (destination *Destination) SetFitV(left float32) error {
	C.HPDF_Destination_SetFitV(destination.destination, C.HPDF_REAL(left))

	return destination.page.pdf.GetLastError()
}

func (destination *Destination) SetFitR(left, bottom, right, top float32) error {
	C.HPDF_Destination_SetFitR(
		destination.destination,
		C.HPDF_REAL(left), C.HPDF_REAL(bottom), C.HPDF_REAL(right), C.HPDF_REAL(top),
	)

	return destination.page.pdf.GetLastError()
}

func (destination *Destination) SetFitB() error {
	C.HPDF_Destination_SetFitB(destination.destination)

	return destination.page.pdf.GetLastError()
}

func (destination *Destination) SetFitBH(top float32) error {
	C.HPDF_Destination_SetFitBH(destination.destination, C.HPDF_REAL(top))

	return destination.page.pdf.GetLastError()
}

func (destination *Destination) SetFitBV(left float32) error {
	C.HPDF_Destination_SetFitBV(destination.destination, C.HPDF_REAL(left))

	return destination.page.pdf.GetLastError()
}
