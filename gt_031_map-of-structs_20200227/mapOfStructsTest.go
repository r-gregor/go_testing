// filename: mapOfStructsTest.go

package main

import "fmt"

type Person struct {
	name 	string
	age 	int
}

type People map[string]*Person

func main() {

	
	var people = make(People) 	// OR: var people People = map[string]*Person{}

	people["Gregor"] = &Person{name: "Gregor Redelongi", age: 52}
	people["Tadeja"] = &Person{name: "Tadeja Mali Redelonghi", age: 50}
	people["Zala"] = &Person{name: "Zala Redelonghi", age: 24}
	people["Mark"] = &Person{name: "Mark Redelonghi", age: 21}
	people["Spela"] = &Person{name: "Špela Redelonghi", age: 11}
	
	fmt.Println("Initial map:")
	people.PrintAll()

	// people["Gregor"].name = "Gregec Navihani"
	NameToChange1 := "Gregor"
	people.ChangeName(NameToChange1, "Gregec Navihani")
	people.PrintPerson(NameToChange1)

	people.ChangeName("Zala", "Zalč Redelonghijus")
	people.PrintPerson("Zala")

	people.ChangeAge("Spela", 12)
	people.PrintPerson("Spela")

	fmt.Println("Final map:")
	people.PrintAll()
}


func (ppl People) PrintAll() {
	for _, p := range ppl {
		fmt.Printf("Name: %s, age: %d\n", p.name, p.age)
	}
	fmt.Println()
}

func (ppl People) PrintPerson(name string) {
	fmt.Printf("Name: %s, age: %d\n", ppl[name].name, ppl[name].age)
	fmt.Println()
}

func (ppl People) ChangeName(name, newName string) {
	fmt.Printf("Changing name \"%s\" to \"%s\" ...\n", ppl[name].name, newName)
	ppl[name].name = newName
}

func (ppl People) ChangeAge(name string, newAge int) {
	fmt.Printf("Changing age of \"%s\" from %d to %d ...\n", ppl[name].name, ppl[name].age, newAge)
	ppl[name].age = newAge
}

