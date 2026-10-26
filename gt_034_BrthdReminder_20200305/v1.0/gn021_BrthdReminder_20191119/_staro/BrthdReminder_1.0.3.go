package main

// v1.0.0    (20191108/1/d -- ORIGINAL)
// v1.0.1   (20191108/2/en)
// v1.0.2   (20191108/3/en)
// v1.0.3   (20191108/4/en -- structs)

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "time"
    "strconv"
)


type person struct {
    nm string
    rd string
    dt string
    yrs int
    dys int
}

var danes = time.Now()
var seznam = make(map[string]string)
var dni int

var seznam2 []person

var st int = 30 + 15 + 15 + 5 + 5
func crta(n int) {
    for i := 0; i < n; i++ {
        fmt.Print("-")
    }
    fmt.Println()
}


func fillData() {

    var prsd string
    var prsyrs int
    var prsdni int
    
    fjl := "ROJSTNIDNEVI.txt"
    file, err := os.Open(fjl)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    
    for scanner.Scan() {
        krneki := strings.Split(scanner.Text(), ",")

        item := new(person)
        item.nm = krneki[0]
        item.rd = krneki[1]

        thisY := time.Time.Format(danes, "2006")
        myRDs := item.rd
        myRD := strings.Split(myRDs, ".")
        myRDThisYs := fmt.Sprintf("%s.%s.%s", myRD[0], myRD[1], thisY)
        myRDThisY, _ := time.Parse("02.01.2006", myRDThisYs)
        
        var YY string
        if danes.After(myRDThisY) {
            danes2 := danes.AddDate(1, 0, 0)
            YY = time.Time.Format(danes2, "2006")
        } else {
            YY = time.Time.Format(danes, "2006")
        }
        
        myRDThisYs = fmt.Sprintf("%s.%s.%s", myRD[0], myRD[1], YY)
        prsd = myRDThisYs
        
        myRDfinal, _ := time.Parse("02.01.2006", myRDThisYs)
        
        diff := myRDfinal.Sub(danes)
        dni = int(diff.Hours() / 24)
        
        now, _      := strconv.Atoi(YY)
        then, _     := strconv.Atoi(myRD[2])
        
        prsyrs = now - then
        prsdni = dni
    
        item.dt = prsd
        item.yrs = prsyrs
        item.dys = prsdni

        // test
        seznam2 = append(seznam2, *item)

    }
}

func DisplayAll() {
    crta(st)
    fmt.Printf("%-30s%-15s%-15s%-5s%-5s\n", "Ime", "RD", "Datum", "Let", "Dni")
    crta(st)
    for _, item := range seznam2 {
       fmt.Printf("%-30s%-15s%-15s%-5d%-5d\n", item.nm, item.rd, item.dt, item.yrs, item.dys)
    }
    crta(st)
}

func main() {
    fillData()
    
    // test
    // fmt.Println(seznam2)
    
    DisplayAll()

}
