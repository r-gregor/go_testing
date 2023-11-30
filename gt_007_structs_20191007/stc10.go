// from filename: go_structs_20191007.txt
// https://www.golangprograms.com/go-language/struct.html
// stc10.go

/* Find Type of Struct in Go Programming Language
   The reflect package support to check the underlying type of a struct.
*/

package main

import (
    "fmt"
    "reflect"
)

type rectangle struct {
    length  float64
    breadth float64
    color   string
}

func main() {
    var rect1 = rectangle{10, 20, "Green"}
    fmt.Println(reflect.TypeOf(rect1))         // main.rectangle
    fmt.Println(reflect.ValueOf(rect1).Kind()) // struct

    rect2 := rectangle{length: 10, breadth: 20, color: "Green"}
    fmt.Println(reflect.TypeOf(rect2))         // main.rectangle
    fmt.Println(reflect.ValueOf(rect2).Kind()) // struct

    rect3 := new(rectangle)
    fmt.Println(reflect.TypeOf(rect3))         // *main.rectangle
    fmt.Println(reflect.ValueOf(rect3).Kind()) // ptr

    var rect4 = &rectangle{}
    fmt.Println(reflect.TypeOf(rect4))         // *main.rectangle
    fmt.Println(reflect.ValueOf(rect4).Kind()) // ptr
}
