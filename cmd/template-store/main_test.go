package main

import (
	"testing"
)

func TestDirConfirm(t *testing.T) {
	var tf bool
	tf = dirConfirm()
	if tf != false {
		t.Error("\nactual: ", tf, "\nexpect: ", false)
	}

	t.Log("Done TestDirConfirm")
}
