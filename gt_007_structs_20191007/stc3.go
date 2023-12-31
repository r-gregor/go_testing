// from filename: go_structs_20191007.txt
// https://www.golangprograms.com/go-language/struct.html
// stc3.go

/* Creating a Struct Instance Using a Struct Literal
   Creates an instance of rectangle struct by using a struct literal and assigning values to the fields
   of the struct.
*/

package main

import "fmt"

type rectangle struct {
    length  int
    breadth int
    color   string
}

func main() {
    var rect1 = rectangle{10, 20, "Green"}
    fmt.Println(rect1)

    var rect2 = rectangle{length: 10, color: "Green"} // breadth value skipped
    fmt.Println(rect2)

    rect3 := rectangle{10, 20, "Green"}
    fmt.Println(rect3)

    rect4 := rectangle{length: 10, breadth: 20, color: "Green"}
    fmt.Println(rect4)

    rect5 := rectangle{breadth: 20, color: "Green"} // length value skipped
    fmt.Println(rect5)
}
