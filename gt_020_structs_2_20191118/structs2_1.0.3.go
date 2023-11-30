// v1.0.0	(20191117/1/d)	-- ORIGINAL
// v1.0.1	(20191117/2/d)
// v1.0.2	(20191117/3/d)
// v1.0.3	(20191117/4/en)

package main

import "fmt"

type names1 []string

type names2 struct {
	myNames []string
}

func main() {

	seznam2 := names2{[]string{"Tadeja"}}
	fmt.Println("print-1:", seznam2)

	seznam2.Add("Zala")
	seznam2.Add("Špela")
	fmt.Println("print-2:", seznam2)

	fmt.Println("print-3:")
	for _, name2 := range seznam2.myNames {
		fmt.Printf("Ime: %s\n", name2)
	}

	// seznam1 := names1{}
	seznam1 := names1{"Oči"}
	seznam1 = append(seznam1, "Mark")
	fmt.Println("print-4:", seznam1)

	seznam1.Add("Gregor")
	fmt.Println("print-5:", seznam1)

	seznam1.Add2("SiSi")
	fmt.Println("print-6:", seznam1)
}

func (m1 *names2) Add(el2 string) {
	m1.myNames = append(m1.myNames, el2)
	fmt.Printf("Adding %s ...\n", el2)
}

func (m1 *names1) Add(el1 string) {
	*m1 = append(*m1, el1)
	fmt.Printf("Adding %s ...\n", el1)
	fmt.Println("Inside function:", *m1)
}

// Add2	-- no pointer --> result only in function
func (m1 names1) Add2(el2 string) {
	m1 = append(m1, el2)
	fmt.Printf("Adding %s ...\n", el2)
	fmt.Println("Inside function:", m1)
}
