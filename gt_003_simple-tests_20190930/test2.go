// test2.go

package main

import "fmt"

var pt1 float64 = 300.00
var pt2 float64 = 301.15
var razdalja float64 = 100.00
var padec float64



func main() {
	padec = (pt2 - pt1) / razdalja
	fmt.Printf("Točka 1: %.2f\n", pt1)
	fmt.Printf("Točka 2: %.2f\n", pt2)
	fmt.Printf("Razdalja: %.2f\n", razdalja)
	fmt.Println("==============================")
	fmt.Printf("Padec: %.2f %%\n", padec * 100)
	


}

