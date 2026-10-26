// test3.go

package main

import "fmt"
import "math"
import "number"

func main() {

	for i := 0; i < 101; i++ {
		//
		base := 2.0
		exp := float64(i)
		var result float64 = math.Pow(base, exp)
		fmt.Printf("%.0f^%-3.0f = %.0f\n", base, exp, result)
	}

	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d  e^x = %8.3f\n", x, math.Exp(float64(x)))
	}


}

