package main

import (
	"fmt"
	"strings"
)

type myslice []int

// MAIN
func main() {

	L1 := myslice{1, 2, 3, 4, 5}
	L1.Info()

	L1.Add(6)
	L1.Info()

	L1.Change(2, 1000)
	L1.Info()

	L1.Change(15, 2000)
	L1.Info()

	L1.Add(600)
	L1.Info()

	L1.ChangeLast(900)
	L1.Info()

	L1.Add(101)
	L1.Info()

	L1.Add(102)
	L1.Info()

	L1.Add(103)
	L1.Info()

	L1.Add(104)
	L1.Info()

	L1.Add(108)
	L1.Info()

	L1.ChangeLast(105)
	L1.Info()

} // end pf MAIN

func (s myslice) Display() {
	for i, v := range s {
		fmt.Println(i, v)
	}
}

func (s myslice) Info() {
	fmt.Printf("Length: %d\nCapacity: %d\nLast element: %d\n", len(s), cap(s), s[len(s)-1])
	fmt.Println(s)
	fmt.Println(strings.Repeat("-", 40))

}

func (s *myslice) Add(el int) {
	fmt.Printf("Trying to append %d ... ", el)
	*s = append(*s, el)
	fmt.Println("Success!")
}

func (s myslice) Change(i int, el int) {
	fmt.Printf("Trying to change index %d to: %d ... ", i, el)
	if i >= len(s) {
		fmt.Printf("Error: Index %d is out of bounaries!\n", i)
	} else {
		s[i] = el
		fmt.Println("Success!")
	}
}

func (s myslice) ChangeLast(el int) {
	last := len(s) - 1
	fmt.Printf("Trying to change index %d to: %d ... ", last, el)
	if last >= len(s) {
		fmt.Printf("Error: Index %d is out of bounaries\n", last)
	} else {
		s[last] = el
		fmt.Println("Success!")
	}
}
