// from filename: go_structs_20191007.txt
// https://www.golangprograms.com/go-language/struct.html
// stc4.go

/* Struct Instantiation using new keyword
   An instance of a struct can also be created with the new keyword. It is then possible to assign data
   values to the data fields using dot notation.
*/

package main

import "fmt"

type rectangle struct {
    length  int
    breadth int
    color   string
}

func main() {
    rect1 := new(rectangle) // rect1 is a pointer to an instance of rectangle
    rect1.length = 10
    rect1.breadth = 20
    rect1.color = "Green"
    fmt.Println(rect1)

    var rect2 = new(rectangle) // rect2 is an instance of rectangle
    rect2.length = 10
    rect2.color = "Red"
    fmt.Println(rect2)
}
