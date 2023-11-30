package main

import (
    "fmt"
    "strings"
)

// custom type
type strslc struct {
    slc []string
}

// custom method
func (s strslc) join(sep string) string {
    return strings.Join(s.slc, sep)
}

// custom method without pointer
func (s strslc) add1(el string) {
    s.slc = append(s.slc, el)
    
    // internal test print
    fmt.Println(s)
}

// custom method with pointer
func (s *strslc) add(el string) {
    s.slc = append(s.slc, el)
    // internal test print
    fmt.Println(s)
}


var besede strslc

func main() {
    besede.slc = append(besede.slc, "ENA")
    besede.slc = append(besede.slc, "DVA")
    besede.slc = append(besede.slc, "TRI")
    besede.slc = append(besede.slc, "ŠTIRI")
    besede.slc = append(besede.slc, "PET")
    
    fmt.Println(besede.join(" + "))
    fmt.Println(besede.join("\n"))
    fmt.Println(besede.join(" ### "))
    fmt.Println(besede.join(", "))
    
    fmt.Println("\n" + "adding TISOČ, DVA TISOČ, TRI TISOČ with pointer to slice:")
    besede.add("TISOČ")
    besede.add("DVA TISOČ")
    besede.add("TRI TISOČ")
    fmt.Println(besede.join(", "))
    
    fmt.Println("\n" + "adding DESET TISOČ, DVAJSET TISOČ withOUT pointer to slice:")
    besede.add1("DESET TISOČ")
    besede.add1("DVAJSET TISOČ")
    fmt.Println(besede.join(", "))
}

/* OUTPUT
$> gregor.redelonghi@cygwin-en [/home/gregor.redelonghi/GO_en_testing]
$> go run custom_methods_20191030.go
ENA + DVA + TRI + ŠTIRI + PET
ENA
DVA
TRI
ŠTIRI
PET
ENA ### DVA ### TRI ### ŠTIRI ### PET
ENA, DVA, TRI, ŠTIRI, PET

adding TISOČ, DVA TISOČ, TRI TISOČ with pointer to slice:
&{[ENA DVA TRI ŠTIRI PET TISOČ]}
&{[ENA DVA TRI ŠTIRI PET TISOČ DVA TISOČ]}
&{[ENA DVA TRI ŠTIRI PET TISOČ DVA TISOČ TRI TISOČ]}
ENA, DVA, TRI, ŠTIRI, PET, TISOČ, DVA TISOČ, TRI TISOČ

adding DESET TISOČ, DVAJSET TISOČ withOUT pointer to slice:
{[ENA DVA TRI ŠTIRI PET TISOČ DVA TISOČ TRI TISOČ DESET TISOČ]}
{[ENA DVA TRI ŠTIRI PET TISOČ DVA TISOČ TRI TISOČ DVAJSET TISOČ]}
ENA, DVA, TRI, ŠTIRI, PET, TISOČ, DVA TISOČ, TRI TISOČ
*/