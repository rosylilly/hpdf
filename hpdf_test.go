package hpdf

import (
	. "testing"
)

func TestVersion(t *T) {
	ver := Version()

	t.Log(ver)
}
