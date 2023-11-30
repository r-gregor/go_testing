package main

import (
	"fmt"
	"strconv"
	"strings"
)

type oseba struct {
	name string
	age  int
}

func main() {

	var seznam []oseba
	seznam = append(seznam, oseba{name: "Gregor", age: 52})
	seznam = append(seznam, oseba{name: "Tadeja", age: 50})
	seznam = append(seznam, oseba{name: "Zala", age: 24})
	seznam = append(seznam, oseba{name: "Mark", age: 21})
	seznam = append(seznam, oseba{name: "Špela", age: 12})

	sum := 0
	var names []string
	for _, el := range seznam {
		expr := el.GetName() + "(" + strconv.Itoa(el.GetAge()) + ")" // e.g. Gregor(52)
		names = append(names, expr)
		sum += el.GetAge()
	}

	famnames := strings.Join(names[:len(names)-1], ", ") + " and " + names[len(names)-1]
	fmt.Print("The cummulative age for family members: ")
	fmt.Print(famnames)
	fmt.Printf(" is ... %d years.\n", sum)

}

func (el *oseba) GetName() string {
	return el.name
}

func (el *oseba) GetAge() int {
	return el.age
}
