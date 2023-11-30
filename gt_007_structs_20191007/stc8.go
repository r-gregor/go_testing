// from filename: go_structs_20191007.txt
// https://www.golangprograms.com/go-language/struct.html
// stc8.go

/* Add Method to Struct Type
   You can also add methods to struct types using a method receiver. A method EmpInfo is added to the
   Employee struct.
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

func (e Employee) EmpInfo() string {
    fmt.Println(e.FirstName, e.LastName)
    fmt.Println(e.Age)
    fmt.Println(e.Email)
    for _, info := range e.MonthlySalary {
        fmt.Println("===================")
        fmt.Println(info.Basic)
        fmt.Println(info.HRA)
        fmt.Println(info.TA)
    }
    return "----------------------"
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

    fmt.Println(e.EmpInfo())
}
