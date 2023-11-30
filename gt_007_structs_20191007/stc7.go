// from filename: go_structs_20191007.txt
// https://www.golangprograms.com/go-language/struct.html
// stc7.go

/* Use Field Tags in the Definition of Struct Type
   During the definition of a struct type, optional string values may be added to each field
   declaration.
*/

package main

import (
    "fmt"
    "encoding/json"
)

type Employee struct {
    FirstName  string `json:"firstname"`
    LastName   string `json:"lastname"`
    City string `json:"city"`
}

func main() {
    json_string := `
    {
        "firstname": "Rocky",
        "lastname": "Sting",
        "city": "London"
    }`

    emp1 := new(Employee)
    json.Unmarshal([]byte(json_string), emp1)
    fmt.Println(emp1)

    emp2 := new(Employee)
    emp2.FirstName = "Ramesh"
    emp2.LastName = "Soni"
    emp2.City = "Mumbai"
    jsonStr, _ := json.Marshal(emp2)
    fmt.Printf("%s\n", jsonStr)
}
