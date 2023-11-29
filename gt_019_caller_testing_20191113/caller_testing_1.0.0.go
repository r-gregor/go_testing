package main

// v1.0.0	(20191113/d/)	-- ORIGINAL

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

func getRootPath() string {
	_, rootFjl, _, _ := runtime.Caller(0)
	return rootFjl
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func dirExists(dirname string) bool {
	info, err := os.Stat(dirname)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func checkIfExists(fname string) {
	if fileExists(fname) {
		fmt.Printf("File [ %s ] exists!\n", fname)
	} else if dirExists(fname) {
		fmt.Printf("Dir [ %s ] exists!\n", fname)
	} else {
		fmt.Println("NO file or dir!")
	}
}

func main() {
	myFile := "file01.txt"

	fjl := getRootPath()
	fmt.Println(fjl)

	dironly := filepath.Dir(fjl)
	fmt.Println(dironly)

	newFile := path.Join(dironly, myFile)
	fmt.Println(newFile)

	newDir := path.Join(dironly, "new_directory")
	fmt.Println(newDir)

	checkIfExists(newFile)
	checkIfExists(newDir)

	pth3 := path.Join(dironly, "file01.txt")
	// pth4 := strings.ReplaceAll(pth3, "/", "\\")
	fmt.Println(pth3)
	checkIfExists(pth3)         // NOT working if called from outside of containing dir
	checkIfExists("file02.txt") // NOT working if called from outside of containing dir

}
