package hpdf

/*
#cgo LDFLAGS: -lhpdf -lpng -lz
#include "hpdf.h"
*/
import "C"

type DashMode struct {
	Ptn    [8]uint16
	NumPtn uint
	Phase  uint
}

func dashModeFromHPDFDashMode(dashMode C.HPDF_DashMode) *DashMode {
	dm := &DashMode{
		Ptn:    [8]uint16{},
		NumPtn: uint(dashMode.num_ptn),
		Phase:  uint(dashMode.phase),
	}

	for i := uint(0); i < dm.NumPtn; i++ {
		dm.Ptn[i] = uint16(dashMode.ptn[i])
	}

	return dm
}
