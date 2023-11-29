package main

import "fmt"

func main() {
	var num float64 = 12456289752.556
	var num2 float64 = 124.556
	fmt.Printf("%f (f - 12456289752.556)\n", num)
	fmt.Printf("%E (E -12456289752.556)\n", num)
	fmt.Printf("%g (g - 12456289752.556)\n", num) // g --> f, but if big number --> e
	fmt.Printf("%g (g - 124.556)\n", num2)
	fmt.Println()

	num *= 2
	num3 := num2 * 2000000
	num2 *= 2
	fmt.Printf("%f (f - 12456289752.556*2)\n", num)
	fmt.Printf("%E (E - 12456289752.556*2)\n", num)
	fmt.Printf("%g (g - 12456289752.556*2)\n", num)
	fmt.Printf("%g (g - 124.556*2)\n", num2)
	fmt.Printf("%g (g - 124.556*2000000)\n", num3)
}
