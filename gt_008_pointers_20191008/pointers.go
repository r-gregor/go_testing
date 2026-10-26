// pointers.go

package main

import "fmt"

func main() {
    x := 0.1
    z := 0.1
    y := 100.0
    p := &x
    q := p
    
    fmt.Printf("x: %.2f, y: %.2f\n", x, y)
    fmt.Printf("p := &x => p: %v, *p: %.2f\n",p, *p)    
    fmt.Printf("x: %.2f, z: %.2f, x == z? %t\n", x, z, x == z)
    fmt.Printf("x: %.2f, *p: %.2f, *q = %.2f, x == *p? %t\n", x, *p, *q, x == *p)


    
    *p += y
    fmt.Printf("After: *p += y => x: %.2f, y: %.2f, *p: %.2f\n", x, y, *p)
    
    *q += *p
    fmt.Printf("After: *q += *p => x: %.2f, *q: %.2f, *p: %.2f\n", x, *q, *p)

}
