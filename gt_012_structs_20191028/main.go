// main.go

package main

import "fmt"

// struct Family member: Fmm
type Fmm struct {
    name string
    age int
}


// method to print info
func (f Fmm) PrintInfo() {
    fmt.Printf("Name: %s, age: %d\n", f.name, f.age)
}


// MAINF
func main() {
    gg := Fmm{"Gregor", 51}
    gg.PrintInfo()
    
    gg.name = "Grgur 2"
    gg.age = 52
    
    tt := Fmm{"Tadeja", 49}
    zz := Fmm{"Zala", 23}
    mm := Fmm{"Mark", 20}
    sp := Fmm{"Špela", 11}
    
    // slice of all fmems
    prsns := []Fmm{gg, tt, zz, mm, sp}
    
    // iterate over slice and apply PrintInfo() on them
    for _, pre := range(prsns) {
        pre.PrintInfo()
    }
}
