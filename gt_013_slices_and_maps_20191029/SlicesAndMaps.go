// SlicesAndMaps.go

package main

import (
    "fmt"
)

func GetItem(coll map[string]string, key string) string {
    
    if coll[key] == "" {
        return "No such element in our table!"
    } else {
        return coll[key]
    }
}

func testKey(keylist []string, coll map[string]string) {
    fmt.Println()
    for _, mykey := range(keylist) {
        fmt.Printf("The element with sign \"" + mykey + "\" is: %s.\n", GetItem(coll, mykey))
    }
}


func main() {

    // slices
    slc1 := []int{1, 2, 3}
    fmt.Println(slc1)

    slc2 := append(slc1, 4, 5)
    fmt.Println(slc2)

    slc1 = append(slc1, 10, 9, 8)
    fmt.Println(slc1)

    n1 := 2
    n2 := 1
    fmt.Println(slc1[n1 : len(slc1)-n2])
    
    slc3 := []string{}
    fmt.Println(slc3)
    
    slc3 = append(slc3, "one", "two", "three")
    fmt.Println(slc3)

    // maps
    elements := map[string]string{
      "H":  "Hydrogen",
      "He": "Helium",
      "Li": "Lithium",
      "Be": "Beryllium",
      "B":  "Boron",
      "C":  "Carbon",
      "N":  "Nitrogen",
      "O":  "Oxygen",
      "F":  "Fluorine",
      "Ne": "Neon",
    }
    
    elements2 := map[string]string{
      "Pt":  "Platynum",
      "Au": "*** Gold ***",
      "Ag": "Silver",
      "U": "Uranium",
      "C":  "Carbon",
      "N":  "Nitrogen",
      "O":  "Oxygen",
      "F":  "Fluorine",
      "Ne": "Neon",
    }
    
    
    signs := []string{"H", "He", "U", "C", "N", "Au", "Ag", "O", "Ne", "Pt"}
    
    testKey(signs, elements)
    testKey(signs, elements2)
}
