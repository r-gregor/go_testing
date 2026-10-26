// from filename: go_structs_20191007.txt
// https://www.golangprograms.com/go-language/struct.html
// stc9.go

/* Assign Default Value for Struct Field
   Method of assigning a custom default value can be achieve by using constructor function. Instead of
   creating a struct directly, the Info function can be used to create an Employee struct with a custom
   default value for the Name and Age field.
*/

package main

import "fmt"

type Employee struct {
    Name string
    Age  int
}

func (obj *Employee) Info() {
    if obj.Name == "" {
        obj.Name = "John Doe"
    }
    if obj.Age == 0 {
        obj.Age = 25
    }
}

func main() {
    emp1 := Employee{Name: "Mr. Fred"}
    emp1.Info()
    fmt.Println(emp1)

    emp2 := Employee{Age: 26}
    emp2.Info()
    fmt.Println(emp2)
}

