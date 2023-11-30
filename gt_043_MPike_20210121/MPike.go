package main

import (
	"fmt"
	"os"
	"strconv"
)

type Razred struct {
	razred int
	znesek int
}

var znesek int = 0
var razlika int = 0

func usage() {
	fmt.Print(` 
	Usage: MPike[.exe] [št. pik]

`)
}

func info(pike int, R Razred, vkratnik int, nznesek int) {
	fmt.Printf("razred: %d, vrednost: %d EUR\n", R.razred, R.znesek)
	fmt.Printf("Pike: %d\n", pike)
	if vkratnik != 0 {
		fmt.Printf("razlika: %d - (%d * %d) = %d\n", pike, vkratnik, R.razred, razlika)
		fmt.Printf("dodatek: %d * %d EUR = %d EUR\n", vkratnik, R.znesek, nznesek)
		fmt.Printf("znesek: %d EUR\n", znesek)
		fmt.Println("---")
	} else {
		fmt.Printf("dodatek: %d EUR\n", nznesek)
		fmt.Printf("znesek: %d EUR\n", znesek)
		fmt.Println("---")
	}
}

func izracunPik(pike int, R Razred) {
	vkratnik := pike / R.razred
	ostanek := pike % R.razred
	var nznesek int
	if vkratnik >= 1 {
		nznesek = vkratnik * R.znesek
		znesek += nznesek
		razlika = ostanek
	} else {
		nznesek = 0
		razlika = pike
	}
	info(pike, R, vkratnik, nznesek)
}

func koncniZnesek(pike int) {
	R1 := Razred{250, 5}
	R2 := Razred{1250, 50}
	R3 := Razred{3500, 210}

	izracunPik(pike, R3)
	izracunPik(razlika, R2)
	izracunPik(razlika, R1)
	fmt.Printf("Končni znesek: %d EUR\n", znesek)
}

// MAIN
func main() {

	if len(os.Args) != 2 {
		usage()
		os.Exit(1)

	} else {
		pike, _ := strconv.Atoi(os.Args[1])
		koncniZnesek(pike)
	}
}
