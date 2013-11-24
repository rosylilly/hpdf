package hpdf

/*
#cgo LDFLAGS: -lhpdf -lpng -lz
#include "hpdf.h"
*/
import "C"
import "time"

func timeToHPDFDate(t time.Time) C.HPDF_Date {
	var date C.HPDF_Date

	// Force UTC Time
	t = t.UTC()

	date.year = C.HPDF_INT(t.Year())
	date.month = C.HPDF_INT(t.Month())
	date.day = C.HPDF_INT(t.Day())
	date.hour = C.HPDF_INT(t.Hour())
	date.minutes = C.HPDF_INT(t.Minute())
	date.seconds = C.HPDF_INT(t.Second())
	date.ind = C.char(' ')

	return date
}
