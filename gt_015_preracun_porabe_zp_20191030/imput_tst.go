// input_tst.go

package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
)

// var imput string

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Vnesi vršno porabo in število: [OGV nOGV]:")
    imput, _ := reader.ReadString('\n')
    s := strings.Fields(imput)
    ogv, _ := strconv.ParseFloat(s[0], 64)
    nogv, _ := strconv.Atoi(s[1])
    fmt.Printf("Vršna poraba - OGV:%.3f\n" +
    "število - nOGV: %d\n", ogv, nogv)
    
    skup := ogv * float64(nogv)
    fmt.Printf("Skupna poraba: %.3f", skup)
}
