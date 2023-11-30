// filename: ZpHitrostPlina.go

package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

var msg string = `
 Uporaba:
 $> ZpHitrostPlina [Qd] [Dz] [S]

	Qd ... dejanski pretok [m3/h]
	Dz ... zunanji premer cevi [mm]
	S  ... debelina stene cevi [mm]
`

func izracunHitrostiPlina(Qd float64, Dz float64, S float64) {
	Dn := (Dz - 2*S) / 1000 // [m]

	w := (Qd * 4) / (3.1415927 * (math.Pow(Dn, 2)) * 3600)

	// fmt.Println("\nIzračun dejanskega pretoka zemeljskega plina")
	fmt.Println("\n" + "Vhodni podatki:")
	fmt.Printf("%-25s"+"%8.2f [m3/h]\n", "Dejanski pretok Qd:", Qd)
	fmt.Printf("%-25s"+"%8.2f [mm]\n", "Premer cevi Dz:", Dz)
	fmt.Printf("%-25s"+"%8.2f [mm]\n", "Debelna stene S:", S)

	fmt.Println("-----------------------------------------")

	fmt.Println("Izračun:")
	fmt.Printf("%-25s"+"%8.2f [m/s]\n", "Hitrost plina w:", w)
	fmt.Println("=========================================")
}

func main() {

	if len(os.Args) != 4 {
		fmt.Println(msg)
		os.Exit(0)
	}

	Qd, _ := strconv.ParseFloat(os.Args[1], 64)
	Dz, _ := strconv.ParseFloat(os.Args[2], 64)
	S, _ := strconv.ParseFloat(os.Args[3], 64)

	fmt.Println("\nIzračun hitrosti zemeljskega plina")
	izracunHitrostiPlina(Qd, Dz, S)

}
