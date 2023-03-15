package main

import (
	"fmt"
	"path/filepath"
)

func dirConfirm() bool {
	_, e := filepath.Abs("~/.template")
	if e != nil {
		return false
	}

	return true
}

func main() {
	fmt.Println("vim-go")
}
