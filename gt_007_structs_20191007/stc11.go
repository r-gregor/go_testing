// from filename: go_structs_20191007.txt
// https://www.golangprograms.com/go-language/struct.html
// stc11.go

/* Comparing Structs with the Different Values Assigned to Data Fields
   Structs of the same type can be compared using comparison operator.
*/

package main

import "fmt"

type rectangle struct {
    length  float64
    breadth float64
    color   string
}

func main() {
    var rect1 = rectangle{10, 20, "Green"}
    rect2 := rectangle{length: 20, breadth: 10, color: "Red"}

    if rect1 == rect2 {
        fmt.Println("True")
    } else {
        fmt.Println("False")
    }

    rect3 := new(rectangle)
    var rect4 = &rectangle{}

    if rect3 == rect4 {
        fmt.Println("True")
    } else {
        fmt.Println("False")
    }
}
