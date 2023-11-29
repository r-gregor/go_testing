package main

// v1.0.0	(20191128/1/en)	-- ORIGINAL

import (
	"fmt"
	"strings"
)

var id int64 = 0
var members = []person{}

type person struct {
	myID   int64
	name   string
	gender string
}

func main() {
	addPerson("Gregor", "M")
	addPerson("Tadeja", "F")
	addPerson("Špela", "F")
	addPerson("Mark", "M")
	addPerson("tAdeja", "F")
	addPerson("Sisi", "F")
	addPerson("Arnold", "M")
	addPerson("Bertold", "M")
	addPerson("Umberto", "M")
	addPerson("ŠPELA", "F")
	addPerson("mark", "M")
	addPerson("Alexander", "M")
	addPerson("Michael", "M")
	addPerson("GrEgOr", "M")
	addPerson("Nick", "M")
	addPerson("Krešimir Đurić", "M")

	showMembers()

}

func addPerson(name string, gender string) {
	fmt.Printf("Adding %s to list ", name)

	var exists bool = false
	for _, prs := range members {

		if strings.ToUpper(name) == strings.ToUpper(prs.name) {
			fmt.Printf("... %s already in a list! Next ID: %d\n", name, id+1)
			exists = true
			break
		} else {
			exists = false
		}
	}

	if !exists {
		id++
		var p = person{id, name, gender}
		members = append(members, p)
		fmt.Printf("at ID = %d ... [OK]\n", id)
	}
}

func showMembers() {
	var num int = len(members)

	fmt.Printf("  %d persons in a list:\n", num)
	for _, p := range members {
		fmt.Printf("%6d - %s, %s\n", p.myID, p.name, p.gender)
	}
}
