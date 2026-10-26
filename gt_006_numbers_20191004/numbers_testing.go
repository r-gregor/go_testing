// numbers_testing.go

package main

import (
    "fmt"
    "strconv"
)

func main() {
    num := 1.2345678910111213141516170181920
    fmt.Printf("Number = %e\n", num)
    
    fmt.Println(5 / 2)
    fmt.Println(5 % 2)
    
    num2 := 44.478
    num_st := strconv.Itoa(int(num2))
    fmt.Println("The number is: " + num_st)
    
    // fmt.Println("The number is: " + num2) --> error concatenating string with int

}
