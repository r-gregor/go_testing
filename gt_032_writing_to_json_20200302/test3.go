package main

import (
	"encoding/json"
	"os"
)

type Person struct {
	Name   string
	Age    int
	Emails []string
}

func main() {
	var List []Person
	p1 := Person{
		Name:   "Abu Ashraf Masnun",
		Age:    27,
		Emails: []string{"masnun@gmail.com", "masnun@localstaffingllc.com"},
	}

	p2 := Person{
		Name:   "Conan The Barbarian",
		Age:    32,
		Emails: []string{"conan.the.barbarian@gmail.com", "king.of.the.world@ohmaygod.com"},
	}

	List = append(List, p1)
	List = append(List, p2)

	fileWriter, _ := os.Create("output3.json")
	json.NewEncoder(fileWriter).Encode(List)
}
