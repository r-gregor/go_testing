package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	usage := `
Uporaba go_argparse.exe:
  -dt float
        -dt, --dt ... delovni (pogonski) tlak [bar]
  -np float
        -np, --np ... normni pretok [Nm3/h]

    `

	var QN float64
	var Pd float64
	var Hlp string

	flag.StringVar(&Hlp, "h", "", "-h ... pomoč/uporaba")
	flag.Float64Var(&QN, "np", 0, "-np ... normni pretok [Nm3/h]")
	flag.Float64Var(&Pd, "dt", 0, "-dt ... delovni (pogonski) tlak [bar]")

	// to test if flags were set
	required := []string{"np", "dt"}
	help := "h"
	flag.Parse()

	// test flags
	seen := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { seen[f.Name] = true })
	for _, req := range required {
		if !seen[req] {
			// or possibly use `log.Fatalf` instead of:
			fmt.Fprintf(os.Stderr, "missing required -%s argument/flag\n", req)
			os.Exit(2) // the same exit code flag.Parse uses
		} else if seen[help] {
			fmt.Println(usage)
		}
	}

	fmt.Println("\nIzračun dejanskega pretoka zemeljskega plina")
	izracunDejanskegaPretoka(QN, Pd)
}

func izracunDejanskegaPretoka(QN float64, Pd float64) {
	T0 := 273.15  // Normna temperatura [K]
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
