package main

import "fmt"


func main() {
    a, b := "A", "B"

    fmt.Printf("Before change: a = %s and b = %s\n", a, b)

    // change
    a,b = b,a
    fmt.Printf("After change: a = %s and b = %s\n", a, b)

}