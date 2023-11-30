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
	masnun := Person{
		Name:   "Abu Ashraf Masnun",
		Age:    27,
		Emails: []string{"masnun@gmail.com", "masnun@localstaffingllc.com"},
	}

	fileWriter, _ := os.Create("output2.json")
	json.NewEncoder(fileWriter).Encode(masnun)
}
