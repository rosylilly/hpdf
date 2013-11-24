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

func TestUseJPEncodings(t *T) {
	pdf, _ := New()

	err := pdf.UseJPEncodings()

	if err != nil {
		t.Fatal(err)
	}
}

func TestUseKREncodings(t *T) {
	pdf, _ := New()

	err := pdf.UseKREncodings()

	if err != nil {
		t.Fatal(err)
	}
}

func TestUseCNTEncodings(t *T) {
	pdf, _ := New()

	err := pdf.UseCNTEncodings()

	if err != nil {
		t.Fatal(err)
	}
}

func TestUseCNSEncodings(t *T) {
	pdf, _ := New()

	err := pdf.UseCNSEncodings()

	if err != nil {
		t.Fatal(err)
	}
}

func TestUseUTFEncodings(t *T) {
	pdf, _ := New()

	err := pdf.UseUTFEncodings()

	if err != nil && err.Error() != "UTF Encoding is not supported" {
		t.Fatal(err)
	}
}

func TestEncoderGetType(t *T) {
	pdf, _ := New()

	err := pdf.SetCurrentEncoder("StandardEncoding")

	if err != nil {
		t.Fatal(err)
	}

	encoder := pdf.GetCurrentEncoder()

	if encoder == nil {
		t.Fatal(pdf.GetLastError())
	}

	encoderType := encoder.GetType()

	if encoderType != ENCODER_TYPE_SINGLE_BYTE {
		t.Fatal(pdf.GetLastError())
	}
}

func TestEncoderGetByteType(t *T) {
	pdf, _ := New()

	err := pdf.SetCurrentEncoder("StandardEncoding")

	if err != nil {
		t.Fatal(err)
	}

	encoder := pdf.GetCurrentEncoder()

	if encoder == nil {
		t.Fatal(pdf.GetLastError())
	}

	byteType := encoder.GetByteType("a", 0)

	if byteType != BYTE_TYPE_SINGLE {
		t.Fatal(pdf.GetLastError())
	}
}

func TestEncoderGetUnicode(t *T) {
	pdf, _ := New()

	err := pdf.SetCurrentEncoder("StandardEncoding")

	if err != nil {
		t.Fatal(err)
	}

	encoder := pdf.GetCurrentEncoder()

	if encoder == nil {
		t.Fatal(pdf.GetLastError())
	}

	unicode := encoder.GetUnicode(97)

	if unicode != 'a' {
		t.Fatal(pdf.GetLastError())
	}
}

func TestEncoderGetWritingMode(t *T) {
	pdf, _ := New()

	err := pdf.SetCurrentEncoder("StandardEncoding")

	if err != nil {
		t.Fatal(err)
	}

	encoder := pdf.GetCurrentEncoder()

	if encoder == nil {
		t.Fatal(pdf.GetLastError())
	}

	writingMode := encoder.GetWritingMode()

	if writingMode != WMODE_HORIZONTAL {
		t.Fatal(pdf.GetLastError())
	}
}
