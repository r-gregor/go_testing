// from filename: go_structs_20191007.txt
// https://www.golangprograms.com/go-language/struct.html
// stc12.go

/* Copy Struct Type Using Value and Pointer Reference
   r2 will be the same as r1, it is a copy of r1 rather than a reference to it. Any changes made to r2
   will not be applied to r1 and vice versa. When r3 is updated, the underlying memory that is assigned
   to r1 is updated.
*/

package main

import "fmt"

type rectangle struct {
    length  float64
    breadth float64
    color   string
}

func main() {
    r1 := rectangle{10, 20, "Green"}
    fmt.Println(r1)

    r2 := r1
    r2.color = "Pink"
    fmt.Println(r2)

    r3 := &r1
    r3.color = "Red"
    fmt.Println(r3)

    fmt.Println(r1)
}
