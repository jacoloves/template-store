package main

import (
	"os"
	"testing"
)

func TestDirConfirm(t *testing.T) {
	homedir, _ := os.UserHomeDir()
	templateDir := homedir + "/.template"

	os.MkdirAll(templateDir, 0755)
	if !dirConfirm() {
		t.Error("not exist template directory")
	}

	os.RemoveAll(templateDir)
	if !dirConfirm() {
		t.Error("exist template directory")
	}
}
