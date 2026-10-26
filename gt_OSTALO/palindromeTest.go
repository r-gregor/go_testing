package main

import "fmt"

func main() {
    var S = []string{"ana Voli MiloVana", "Gregor Redelonghi", "ANA", "MIHA", "ABBA", "AC/DC", "AABBcBBAA"}

    for _, str := range(S) {
        fmt.Print(str + " --> ")
        i := 0
        j := len(str) - 1

        for {
            L := str[i]
            D := str[j]
            if L != D {
                fmt.Println("NOT palindrom!")
            break
            }

            if i < j {
                fmt.Println("YES palindrom!")
            break

            i++
            j--
            }
        }
    }
}
