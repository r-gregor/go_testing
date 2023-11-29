// pnt.go

package main

import "fmt"

func change(a *string) {
	fmt.Printf("Original value: \"%v\"\n", *a)
	*a = "Changed!"
	fmt.Printf("Original value changed to: \"%v\" inside change() func\n", *a)
}

func noChange(a string) {
	fmt.Printf("Original value: \"%s\"\n", a)
	a = "Changed!"
	fmt.Printf("Original value changed to: \"%v\" inside noChange() func\n", a)
}

func main() {
	var x int = 1500
	fmt.Println("Variable x value:", x)
	fmt.Println("Variable x address (&x):", &x)

	y := &x
	fmt.Println("Variable y := &x value:", y)
	fmt.Println("Variable *y value BEFORE change:", *y)

	*y = 300
	fmt.Println("Variable *y value AFTER change:", *y)
	fmt.Println("Variable x value AFTER change:", x)

	fmt.Println()
	var word string = "Original"
	change(&word)
	fmt.Printf("Original value after the call to change() func: \"%v\"\n", word)

	fmt.Println()
	var word2 string = "Original"
	noChange(word2)
	fmt.Printf("Original value after the call to noChange() func: \"%v\"\n", word2)

}
