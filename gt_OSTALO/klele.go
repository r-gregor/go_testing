package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	bin, _ := os.Executable()
	pth, _ := filepath.Abs(filepath.Dir(bin))
	fmt.Println(pth)
	fname := "ROJSTNIDNEVI.txt"

	fpath := filepath.Join(pth, fname)
	fmt.Println(fpath)
}
