// timeAndDate.go

package main

import (
    "fmt"
    "time"
)

// function PrintOut uses two parameters:
//      1 - an array of format strings, and
//      2 - a variable containing current timr from main
func PrintOut(fmtArr []string, when time.Time) {

    fmt.Println("\n" + "func PrintOut(fmtArr []string, when time.Time)")
    fmt.Println("Formated Date/Time combos: (based on \"Mon Jan 2 15:04:05 MST 2006\" or \"01/02 2006 03:04:05PM '06 -0700\"")
    fmt.Println("------------------------------------------------------")
    fmt.Printf("%-28s %s\n", "Format", "Result")
    fmt.Println("------------------------------------------------------")
    for i := 0; i < len(fmtArr); i++ {
        fmt.Printf("%-28v %s\n", fmtArr[i] + ":", when.Format(fmtArr[i]))
    }
    fmt.Println("------------------------------------------------------")
}


// function PrintOut2 uses ONLY ONE parameter:
//      1 - an array of format strings
func PrintOut2(fmtArr []string) {

    fmt.Println("\n" + "PrintOut2(fmtArr []string)")
    fmt.Println("Formated Date/Time combos: (based on \"Mon Jan 2 15:04:05 MST 2006\" or \"01/02 2006 03:04:05PM '06 -0700\"")
    fmt.Println("------------------------------------------------------")
    fmt.Printf("%-28s %s\n", "Format", "Result")
    fmt.Println("------------------------------------------------------")
    for i := 0; i < len(fmtArr); i++ {
        fmt.Printf("%-28v %s\n", fmtArr[i] + ":", time.Now().Format(fmtArr[i]))
    }
    fmt.Println("------------------------------------------------------")
}

func main() {
    zdaj := time.Now()
    fmt.Println("Current time: ", zdaj)
    
    // format in (Mon Jan 2 15:04:05 MST 2006) or (01/02 2006 03:04:05PM '06 -0700)
    var fmts = []string{"02-Jan-2006", "20060102", "20060102_0304", "20060102_030405", "Mon, 02/01-2006, 03:04:05", "Mon, 02. Jan, 2006", "02/01/2006", "02.01.2006"}
    
    PrintOut(fmts, zdaj)
    
    PrintOut2(fmts)
}