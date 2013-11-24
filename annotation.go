package hpdf

/*
#cgo LDFLAGS: -lhpdf -lpng -lz
#include "hpdf.h"
*/
import "C"
import (
	"unsafe"
)

func (page *Page) CreateTextAnnot(
	rect *Rect, text string, encoder *Encoder,
) (*TextAnnotation, error) {
	ctext := C.CString(text)
	var cencoder C.HPDF_Encoder = nil
	if encoder != nil {
		cencoder = encoder.encoder
	}

	annotation := C.HPDF_Page_CreateTextAnnot(
		page.page, rect.toHpdf(), ctext, cencoder,
	)
	C.free(unsafe.Pointer(ctext))

	if annotation != nil {
		return &TextAnnotation{Annotation{annotation, page.pdf}}, nil
	} else {
		return nil, page.pdf.GetLastError()
	}
}

func (page *Page) CreateLinkAnnot(
	rect *Rect, destination *Destination,
) (*LinkAnnotation, error) {
	annotation := C.HPDF_Page_CreateLinkAnnot(
		page.page, rect.toHpdf(), destination.destination,
	)

	if annotation != nil {
		return &LinkAnnotation{Annotation{annotation, page.pdf}}, nil
	} else {
		return nil, page.pdf.GetLastError()
	}
}

func (page *Page) CreateURILinkAnnot(
	rect *Rect, uri string,
) (*URILinkAnnotation, error) {
	curi := C.CString(uri)

	annotation := C.HPDF_Page_CreateURILinkAnnot(
		page.page, rect.toHpdf(), curi,
	)
	C.free(unsafe.Pointer(curi))

	if annotation != nil {
		return &URILinkAnnotation{Annotation{annotation, page.pdf}}, nil
	} else {
		return nil, page.pdf.GetLastError()
	}
}

type Annotation struct {
	annotation C.HPDF_Annotation
	pdf        *PDF
}

func (annotation *Annotation) SetBorderStyle(
	subtype BSSubtype, width float32, dashOn, dashOff, dashPhase uint16,
) error {
	C.HPDF_Annotation_SetBorderStyle(
		annotation.annotation,
		C.HPDF_BSSubtype(subtype),
		C.HPDF_REAL(width),
		C.HPDF_UINT16(dashOn), C.HPDF_UINT16(dashOff), C.HPDF_UINT16(dashPhase),
	)
	return annotation.pdf.GetLastError()
}

type TextAnnotation struct {
	Annotation
}

func (annotation *TextAnnotation) SetIcon(icon AnnotIcon) error {
	C.HPDF_TextAnnot_SetIcon(
		annotation.annotation, C.HPDF_AnnotIcon(icon),
	)
	return annotation.pdf.GetLastError()
}

func (annotation *TextAnnotation) SetOpened(opened bool) error {
	var copend C.HPDF_BOOL = C.HPDF_FALSE
	if opened {
		copend = C.HPDF_TRUE
	}
	C.HPDF_TextAnnot_SetOpened(
		annotation.annotation, copend,
	)
	return annotation.pdf.GetLastError()
}

type LinkAnnotation struct {
	Annotation
}

func (annotation *LinkAnnotation) SetHighlightMode(
	annotHightlightMode AnnotHighlightMode,
) error {
	C.HPDF_LinkAnnot_SetHighlightMode(
		annotation.annotation, C.HPDF_AnnotHighlightMode(annotHightlightMode),
	)
	return annotation.pdf.GetLastError()
}

func (annotation *LinkAnnotation) SetBorderStyle(
	width float32, dashOn, dashOff uint16,
) error {
	C.HPDF_LinkAnnot_SetBorderStyle(
		annotation.annotation,
		C.HPDF_REAL(width), C.HPDF_UINT16(dashOn), C.HPDF_UINT16(dashOff),
	)
	return annotation.pdf.GetLastError()
}

type URILinkAnnotation struct {
	Annotation
}
