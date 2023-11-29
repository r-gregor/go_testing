package main

// v1.0.0	(20191120/1/en)		-- ORIGINAL

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	binBybin2()
	binByOne2()
}

func binByOne() {
	for i := 0; i < 257; i++ {
		if i%10 == 0 {
			fmt.Printf("\n%-5d:%-9b| ", i, i)
		} else {
			fmt.Printf("%-5d:%-9b| ", i, i)
		}
	}
	fmt.Println()
}

func binByOne2() {
	for i := 0; i < 257; i++ {
		if i%10 == 0 {
			fmt.Printf("\n%-5d:%-9s| ", i, strconv.FormatInt(int64(i), 2))
		} else {
			fmt.Printf("%-5d:%-9s| ", i, strconv.FormatInt(int64(i), 2))
		}
	}
	fmt.Println()
}

func binBybin() {
	fmt.Printf("%-4s%-6s%-15s\n", "i", "2^i", "BIN")
	for i := 0; i < 11; i++ {
		num := math.Pow(2, float64(i))
		fmt.Printf("%-4d%-6d%-15b\n", int(i), int(num), int(num))
	}
	fmt.Println()
}

func binBybin2() {
	fmt.Printf("%-4s%-6s%-15s\n", "i", "2^i", "BIN")
	for i := 0; i < 11; i++ {
		num := math.Pow(2, float64(i))
		fmt.Printf("%-4d%-6d%-15s\n", i, int(num), strconv.FormatInt(int64(num), 2))
	}
	fmt.Println()
}
