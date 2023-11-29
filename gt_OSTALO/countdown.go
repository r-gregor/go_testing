package main

import "fmt"

func main() {

	// V1
	S := ""
	for i := 0; i < 10; i++ {
		S += "G"
		fmt.Println(S)
	}

	// V2
	for j := 0; j < 10; j++ {
		for k := 0; k < (j + 1); k++ {
			fmt.Print("G")
		}
		fmt.Print("\n")
	}

	// V3
	for j := 1; j < 11; j++ {
		for k := 1; k < (j + 1); k++ {
			fmt.Print("G")
		}
		fmt.Print("\n")
	}
}
