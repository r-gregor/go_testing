// v1.0.0.	(20191117/1/d)	-- ORIGINAL

package main

import "fmt"

type myList []string

type imena2 struct {
	myname []string
}

func main() {
	fmt.Printf("filename: %s\n", "structs2_1.0.0.go")

	imena := myList{}
	imena.Add("Gregor")
	imena.Add("Zala")
	imena.Add("Nobady")

	seznam2 := imena2{}
	seznam2.Add("Tadeja")
	seznam2.Add("Mark")

	fmt.Println(imena)
	fmt.Println(seznam2)

	for _, name := range imena {
		fmt.Printf("Ime: %s\n", name)
	}

	for _, name2 := range seznam2.myname {
		fmt.Printf("Ime: %s\n", name2)
	}
}

func (ll *myList) Add(el string) {
	(*ll) = append((*ll), el)
}

func (m *imena2) Add(el string) {
	(*m).myname = append((*m).myname, el)
}
