// filename: ZpPretokNormniDejanski.go

package main

import (
	"fmt"
	"os"
	"strconv"
)

var msg string = `
 Uporaba:
 $> ZpPretokNormniDejanski [QN] [Pd]

	QN ... normni pretok [Nm3/h]
	Pd ... delovni tlak  [bar]
`

func izracunDejanskegaPretoka(QN float64, Pd float64) {
	T0 := 288.15  // Normna temperatura [K]
	P0 := 1.01325 // Normni tlak [bar]
	var Patm float64
	// var Tz float64 = 279.15 // Računska temperatura - zunaj  (6°C) [K]
	Tn := 288.15 // Računska temperatura - notri (15°C) [K]

	H := 298.0 // povprečna nadmorska višina LJ = 298 m
	var fR float64
	var Qd float64

	Patm = 1016.0 - (0.12 * H)
	fR = (T0 / Tn) * (((Patm / 1000) + Pd) / P0)
	Qd = QN / fR

	// fmt.Println("\nIzračun dejanskega pretoka zemeljskega plina")
	fmt.Println("\n" + "Vhodni podatki:")
	fmt.Printf("%-25s"+"%8.2f [bar]\n", "Delovni tlak Pd:", Pd)
	fmt.Printf("%-25s"+"%8.2f [Nm3/h]\n", "Normni pretok QN:", QN)

	fmt.Println("-----------------------------------------")

	fmt.Println("Izračun:")
	fmt.Printf("%-25s"+"%8.0f [mbar]\n", "Atmosferski tlak Patm:", Patm)
	fmt.Printf("%-25s"+"%8.2f [-]\n", "Redukcijski faktor fR:", fR)
	fmt.Printf("%-25s"+"%8.2f [m3/h]\n", "Dejanski pretok Qd:", Qd)
	fmt.Println("=========================================")
}

func main() {

	if len(os.Args) != 3 {
		fmt.Println(msg)
		os.Exit(0)
	}

	QN, _ := strconv.ParseFloat(os.Args[1], 64)
	Pd, _ := strconv.ParseFloat(os.Args[2], 64)

	fmt.Println("\nIzračun dejanskega pretoka zemeljskega plina")
	izracunDejanskegaPretoka(QN, Pd)

}
