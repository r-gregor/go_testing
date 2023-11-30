// v1.0.0	(20191117/1/d)	-- ORIGINAL
// v1.0.1	(20191117/2/d)

package main

import "fmt"

type names1 []string

type names2 struct {
	myNames []string
}

func main() {

	imena := names1{}
	imena.Add("Gregor")
	imena.Add("Zala")
	imena.Add("Nobady")

	seznam2 := names2{}
	seznam2.Add("Tadeja")
	seznam2.Add("Mark")

	AddNameToList("NebukaDnezar", &imena)

	fmt.Println(imena)
	fmt.Println(seznam2)

	for _, name := range imena {
		fmt.Printf("Ime: %s\n", name)
	}

	for _, name2 := range seznam2.myNames {
		fmt.Printf("Ime: %s\n", name2)
	}
}

func (ll *names1) Add(el string) {
	*ll = append(*ll, el)
}

func (m *names2) Add(el string) {
	(*m).myNames = append((*m).myNames, el)
}

// AddNameToList add name to list
func AddNameToList(name string, list *names1) {
	*list = append(*list, name)
}
