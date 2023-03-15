package main

import (
	"flag"
	"os"
)

var regFlag bool

func dirConfirm() bool {
	homedir, _ := os.UserHomeDir()
	templateDir := homedir + "/.template"

	if _, err := os.Stat(templateDir); os.IsNotExist(err) {
		return false
	}

	return true
}

func createTemplateDir() {
	homedir, _ := os.UserHomeDir()
	templateDir := homedir + "/.template"

	os.MkdirAll(templateDir, 0777)
}

func init() {
	flag.BoolVar(&regFlag, "r", false, "Register Flag")
}

func main() {
	existFlg := dirConfirm()

	if !existFlg {
		createTemplateDir()
	}
}
