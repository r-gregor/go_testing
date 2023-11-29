package main

import (
	"fmt"
	"path"
	"path/filepath"
	"runtime"
)

// v1.0.0 (works on WIN10)
// C:\Users\gregor.redelonghi\BRISI\RROOTT\src\scriptAbsPath_20191112\afile.txt
//																	 ^
func main() {

	_, filename, _, _ := runtime.Caller(0)
	absp, err := filepath.Abs(path.Dir(filename))

	fjl := absp + "\\afile.txt"

	if err == nil {
		fmt.Println(fjl)
	}
}

/*
// v1.0.1 (test it on *NIX)
func main() {
	dir, err := filepath.Abs(path.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	// works on win10 --> \\
	// mypath := dir + "\\afile.txt"

	// NOT working on WIN10 --> ..\..\dir/afile.txt
	// C:\Users\gregor.redelonghi\BRISI\RROOTT\src\scriptAbsPath_20191112
	// C:\Users\gregor.redelonghi\BRISI\RROOTT\src\scriptAbsPath_20191112/afile.txt
	//  								                                 ^
	mypath := path.Join(dir, "afile.txt")
	fmt.Println(dir)
	fmt.Println(mypath)
}
*/
