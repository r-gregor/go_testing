// from filename: go_structs_20191007.txt
// https://www.golangprograms.com/go-language/struct.html
// stc6.go

/* Nested Struct Type
   Struct can be nested by creating a Struct type using other Struct types as the type for the fields of
   Struct. Nesting one struct within another can be a useful way to model more complex structures.
*/

package main

import "fmt"

type Salary struct {
        Basic, HRA, TA float64
}

type Employee struct {
        FirstName, LastName, Email string
        Age                        int
        MonthlySalary              []Salary
}

func main() {
        e := Employee{
                FirstName: "Mark",
                LastName:  "Jones",
                Email:     "mark@gmail.com",
                Age:       25,
                MonthlySalary: []Salary{
                        Salary{
                                Basic: 15000.00,
                                HRA:   5000.00,
                                TA:    2000.00,
                        },
                        Salary{
                                Basic: 16000.00,
                                HRA:   5000.00,
                                TA:    2100.00,
                        },
                        Salary{
                                Basic: 17000.00,
                                HRA:   5000.00,
                                TA:    2200.00,
                        },
                },
        }
        fmt.Println(e.FirstName, e.LastName)
        fmt.Println(e.Age)
        fmt.Println(e.Email)
        fmt.Println(e.MonthlySalary[0])
        fmt.Println(e.MonthlySalary[1])
        fmt.Println(e.MonthlySalary[2])
}
