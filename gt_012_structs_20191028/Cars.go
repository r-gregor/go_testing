// Cars.go

package main

import "fmt"

type Car struct {
    TypeName    string
    Model       string
    Year        int
    Family      bool
    Color       string
}

// initial slice of Car structs
var mycars []Car = make([]Car, 0)


func (c Car) GetInfo() string {
    s := fmt.Sprintf(
    "type: %s\n" +
    "model: %s\n" +
    "year: %d\n" +
    "for families: %t\n" +
    "color: %s\n", c.TypeName, c.Model, c.Year, c.Family, c.Color)
    return s
}


func (c Car) GetShortInfo() string {
    s := fmt.Sprintf(
    "%s / " +
    "%s / " +
    "%d / " +
    "family: %t / " +
    "%s\n", c.TypeName, c.Model, c.Year, c.Family, c.Color)
    return s
}

// make a Car object and append it to Cars slice
func MakeCar(TypeName string, Model string, Year int, Family bool, Color string) {
    c := Car{TypeName, Model, Year, Family, Color}
    mycars = append(mycars, c)
}

func main() {
    
    MakeCar("FORD", "S-MAX", 2018, true, "white")
    MakeCar("RENAULT", "Traffic", 1989, false, "blue")
    MakeCar("WV", "Polo", 2014, true, "silver")
    MakeCar("Adria", "Path Finder 8000", 2002, true, "white/green")
    MakeCar("Volvo", "SuperTruck SC1500", 1999, false, "black/red")
    
    // TEST
    // fmt.Println(mycars)

    
    // print all
    for _, car := range mycars {
        fmt.Printf("%s", car.GetShortInfo())
    }
    
    // print only if NOT for families:
    fmt.Println("\n" + "NOT family cars:")
    for _, car := range mycars {
        if car.Family == false {
            fmt.Printf("%s" + "\n", car.GetInfo())
        }
    }
}
