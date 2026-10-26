// from filename: go_structs_20191007.txt
// https://www.golangprograms.com/go-language/struct.html
// strct1.go

/* Declaration of a struct type
   A struct type rectangle is declared that has three data fields of different data-types. Here, struct
   used without instantiate a new instance of that type.
*/


package main

import "fmt"

type rectangle struct {
    length  float64
    breadth float64
    color   string
}

func main() {
    fmt.Println(rectangle{10.5, 25.10, "red"})
}
