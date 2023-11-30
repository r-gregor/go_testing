package main

// v1.0.0	(20191113/d/)	-- ORIGINAL
// V1.0.1	(20191113/en/)	-- added function to deliver absolute path of a file
// 							   getRealPath(filename)

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

func main() {

	checkIfExists("file01.txt")
	checkIfExists("new_directory")
	checkIfExists("subdir1/subsubdir2")
	checkIfExists("file02.txt")
	checkIfExists("media/Movies")

}

func checkIfExists(fname string) {

	if fileExists(fname) {
		fmt.Printf("File [ %s ] exists!\n", fname)
	} else if dirExists(fname) {
		fmt.Printf("Dir [ %s ] exists!\n", fname)
	} else {
		fmt.Printf("NO [ %s ] file or dir!\n", fname)
	}
}

func fileExists(myfname string) bool {

	filename := getRealPath(myfname)
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func dirExists(mydname string) bool {

	dirname := getRealPath(mydname)
	info, err := os.Stat(dirname)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func getRealPath(filename string) string {
	myFile := filename

	fjl := getRootPath()
	// fmt.Println(fjl)

	dironly := filepath.Dir(fjl)
	// fmt.Println(dironly)

	realPath := path.Join(dironly, myFile)
	// fmt.Println(realPath)

	return realPath
}

func getRootPath() string {
	_, rootFjl, _, _ := runtime.Caller(0)
	return rootFjl
}
