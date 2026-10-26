package main

import "fmt"

type IntStrMap map[int]string
type StrStrMap map[string]string

func (m IntStrMap) Add(k int, v string) {
    
    fmt.Printf("Adding %d ==> %s to map ...\n", k, v)
    if m[k] != "" {
        fmt.Printf("\t" + " ...Key %d already exist. Chose another.\n", k)
    } else {
        m[k] = v
    }
    
}

func (m StrStrMap) Add(k string, v string) {
    
    fmt.Printf("Adding %s ==> %s to map ...\n", k, v)
    if m[k] != "" {
        fmt.Printf("\t" + " ...Key %s already exist. Chose another.\n", k)
    } else {
        m[k] = v
    }
    
}

func main() {
    mis := make(IntStrMap)
    mis.Add(1, "ENA")
    mis.Add(2, "DVA")
    mis.Add(3, "TRI")
    mis.Add(1, "štiri")
    mis.Add(4, "ŠTIRI")
    
    fmt.Println("\n" + "Printing map ...")
    fmt.Println(mis)
    
    fmt.Println("----------------------------------------")
    
    mss := make(StrStrMap)
    mss.Add("one", "1")
    mss.Add("two", "2")
    mss.Add("three", "3")
    mss.Add("one", "4")
    mss.Add("four", "4")
    
    fmt.Println("\n" + "Printing map ...")
    fmt.Println(mss)
}
