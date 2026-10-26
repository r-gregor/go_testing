package main

// v1.0.0	(20191120/1/en)		-- ORIGINAL

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	hexByhex2()
	hexByOne2()
}

func hexByOne() {
	for i := 0; i < 257; i++ {
		if i%10 == 0 {
			fmt.Printf("\n%-5d:%-3X| ", i, i)
		} else {
			fmt.Printf("%-5d:%-3X| ", i, i)
		}
	}
	fmt.Println()
}

func hexByOne2() {
	for i := 0; i < 257; i++ {
		if i%10 == 0 {
			fmt.Printf("\n%-5d:%-3s| ", i, strconv.FormatInt(int64(i), 16))
		} else {
			fmt.Printf("%-5d:%-3s| ", i, strconv.FormatInt(int64(i), 16))
		}
	}
	fmt.Println()
}

func hexByhex() {
	fmt.Printf("%-5s%-15s%-15s\n", "i", "16^i", "HEX")
	for i := 0; i < 11; i++ {
		num := math.Pow(16, float64(i))
		fmt.Printf("%-5d%-15d%-15X\n", int(i), int(num), int(num))
	}
	fmt.Println()
}

func hexByhex2() {
	fmt.Printf("%-5s%-15s%-15s\n", "i", "16^i", "HEX")
	for i := 0; i < 11; i++ {
		num := math.Pow(16, float64(i))
		fmt.Printf("%-5d%-15d%-15s\n", i, int(num), strconv.FormatInt(int64(num), 16))
	}
	fmt.Println()
}
