// joining_strings.go

package main

import (
    "fmt"
    "strings"
)

func crt(n int) {
    for i := 0; i < n + 1; i++ {
        fmt.Print("-")
    }
    fmt.Println()
}

func main() {
    L := []string{"Prvi", "Drugi", "Tretji"}

    crt(50) // --------------------
    
    // print without formatting
    fmt.Println(L)
    
    crt(50) // --------------------
    
    // print each el in separate line
    for _, el := range L {
        fmt.Println(el)
    }
    
    crt(50) // --------------------
    
    // print all elements in a row formated (with soaces between)
    S := strings.Join(L, " ")
    fmt.Println(S)
    
    crt(50) // --------------------

}


