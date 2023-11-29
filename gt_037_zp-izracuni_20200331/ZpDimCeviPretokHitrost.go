// filename: ZpDimCeviPretokHitrost.go
package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

// usage
var usage string = `
    Usage:
    ZpDimCeviPretokHitrost [Qd] [maxw]

        Qd ..... dejanski pretok [m3/h]
        maxw ... max hitrost ZP  [m/s]

`

type cev struct {
	Ozn string
	S   []float64
}

func main() {
	cevi := make(map[float64]cev)

	cevi[33.7] = cev{"DN25", []float64{3.25, 3.6}}
	cevi[42.4] = cev{"DN40", []float64{3.25, 3.6}}
	cevi[60.3] = cev{"DN50", []float64{3.2, 3.6, 4.0, 4.5}}
	cevi[88.9] = cev{"DN80", []float64{3.2, 3.6, 4.0, 4.5}}
	cevi[114.3] = cev{"DN100", []float64{3.6, 4.0, 4.5, 5.0}}
	cevi[168.3] = cev{"DN150", []float64{3.6, 4.0, 4.5, 5.0}}
	cevi[219.1] = cev{"DN200", []float64{4.0, 4.5, 5.0, 5.6}}
	cevi[273.0] = cev{"DN250", []float64{4.0, 4.5, 5.0, 5.6}}
	cevi[323.9] = cev{"DN300", []float64{4.0, 4.5, 5.0, 5.6}}

	if len(os.Args) == 1 {
		fmt.Println(usage)

	} else if len(os.Args) == 3 {
		Qd, _ := strconv.ParseFloat(os.Args[1], 64)
		maxw, _ := strconv.ParseFloat(os.Args[2], 64)

		fmt.Printf("Pretok: %.2f m3/h\n", Qd)
		izracun(Qd, maxw, cevi)

	} else {
		fmt.Println("Wrong number of parameters.")
		fmt.Println(usage)
		os.Exit(1)
	}

} // end main

// izračun hitrosti glede na pretok in presek
func hitrost(Qd float64, A float64) float64 {
	w := Qd / (A * 3600)
	return w
}

// izračun preseka glede na dimenzije cevi
func povrsina(Dz float64, S float64) float64 {
	Dn := Dz - 2*S
	Ac := (3.1415927 * math.Pow(Dn, 2)) / 4
	return Ac
}

func izracun(pretok float64, maxw float64, cevi map[float64]cev) {
	Qd := pretok
	razlika := maxw
	zmagovalec := ""

	for k, v := range cevi {
		for _, ds := range v.S {
			A := povrsina(k, ds)
			Am := A / (math.Pow(1000, 2))
			w := hitrost(Qd, Am)

			if w < maxw {
				// izbere tistega znagovalca pri katerem je razlika najmanša!
				if maxw-w < razlika {
					razlika = maxw - w
					// zmagovalec = "--> Izbrana cev: " + str(k) +"x"+ str(ds) + " (" + v[0] + "), hitrost: " + str("{:.3f}".format(w)) + " m/s"
					zmagovalec = fmt.Sprintf(" --> Izbrana cev: %.1fx%.1f (%s), hitrost: %.3f m/s", k, ds, v.Ozn, w)
				}
			} else {
				continue
			} // end if
		} // end for
	} // end for
	S1 := fmt.Sprintf("Max hitrost: %.1f m/s", maxw)
	fmt.Printf("%-22s%s\n", S1, zmagovalec)
}
