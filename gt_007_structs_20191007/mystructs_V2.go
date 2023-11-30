// mystructs_V2.go

package main

import "fmt"

type Person struct {
    name string
    sname string
    addr string
    city string
    zip int
    phone string
}

func (e Person) DisplayInfo() {
    fmt.Printf("%-10s %s %s\n", "Name:", e.name, e.sname)
    fmt.Printf("%-10s %s, %d %s\n", "Address:", e.addr, e.zip, e.city)
    fmt.Printf("%-10s %s\n\n", "Phone:", e.phone)

}


func main() {

    var prs1 = Person{"Gregor", "Redelonghi", "Valvasorjeva ulica 5", "Ljubljana", 1000, "+386 40 88 55 60"}
    
    var prs2 Person
    prs2.name = "Tadeja"
    prs2.sname = "Mali Redelonghi"
    prs2.addr = "Valvasorjeva ulica 5"
    prs2.city = "Ljubljana"
    prs2.zip = 1000
    prs2.phone = "+386 41 629 079"
    
    var prs3 = Person{"Conan", "The Barbarian", "Symeria", "Mist Islands", 1009, "00 904 408 85 560"}

    // array
    // prsns := [...]Person{prs1, prs2, prs3}
    // changed to slice ...
    
    // slice
    // Both work:
    // prsns := make([]Person, 0) --or-- prsns := []Person{}
    prsns := []Person{}
    
    // s = append(s, el1, el2, ...)
    prsns = append(prsns, prs1, prs2, prs3)
    
    
    for _, e := range prsns {
        e.DisplayInfo()
    } 

}

