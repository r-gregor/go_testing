// filename: padec.go

package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {

	var start_p float64
	var end_p float64
	var h_dist float64
	var v_dist float64

	mpadec := 0.003

	if len(os.Args) == 4 {
		start_p, _ = strconv.ParseFloat(os.Args[1], 64)
		end_p, _ = strconv.ParseFloat(os.Args[2], 64)
		h_dist, _ = strconv.ParseFloat(os.Args[3], 64)
		v_dist = start_p - end_p
		ppadec := Padec(start_p, end_p, h_dist)

		fmt.Printf("Začetna točka: %.2f m\n"+
			"Končna točka: %.2f m\n"+
			"Horizontalna razdalja: %.2f m\n"+
			"Vertikalna razdalja: %.2f m\n", start_p, end_p, h_dist, v_dist)
		fmt.Println("=========================================")
		fmt.Printf("Padec: %.3f %%\n"+
			"\n", (ppadec * 100))

		// test
		// fmt.Printf("%.3f %%\n", (ppadec * 100))

		if math.Abs(ppadec) < mpadec {
			fmt.Println("Padec je manjši kot 0.3 %!")
			kv_dist := (mpadec * h_dist)
			fmt.Printf("Višina končne točke za padec vsaj 0.3 %% mora biti:\n"+
				"\t(+) %.3f m ali več.\n", start_p+kv_dist)
			fmt.Printf("\t(-) %.3f m ali manj.\n", start_p-kv_dist)
			fmt.Printf("Vertikalna razdalja mora biti vsaj: (+/-) %.3f m ali več.\n", kv_dist)
		}

	} else {
		Usage()
	}
}

func Usage() {
	fmt.Print(
		`Usage:
padec.exe [pt1] [pt2] [h_razd]
		pt1 ...... začetna višina [m]
		pz2 ...... končna višina [m]
		h_razd ... horizontalna razdalja [m]
`)
}

func Padec(p1 float64, p2 float64, hd float64) float64 {
	return (p1 - p2) / hd
}
