package main

import (
	"os"
)

func dirConfirm() bool {
	homedir, _ := os.UserHomeDir()
	templateDir := homedir + "/.template"

	if _, err := os.Stat(templateDir); os.IsNotExist(err) {
		return false
	}

	return true
}

func main() {
	//dirConfirm()
}
