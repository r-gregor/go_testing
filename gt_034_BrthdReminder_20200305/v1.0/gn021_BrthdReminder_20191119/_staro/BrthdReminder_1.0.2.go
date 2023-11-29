package main

// v1.0.0    (20191108/1/d -- ORIGINAL)
// v1.0.1   (20191108/2/en)
// v1.0.2   (20191108/3/en)

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
var st int = 30 + 15 + 15 + 5 + 5

func crta(n int) {
    for i := 0; i < n; i++ {
        fmt.Print("-")
    }
    fmt.Println()
}

func Display2(mynm string) {
    for name, _ := range seznam {
        if strings.Contains(name, mynm) {
            prsData := seznam[name]
            fmt.Printf("RD for %s: %s\n", name, prsData)
        }
    }
}

func Display(mynm string) {
    names := []string{}
    count := 0
    for name, _ := range seznam {
        if strings.Contains(name, mynm) {
            count += 1
            names = append(names, name)
        }
    }
    
    if count > 0 {
        crta(st)
        fmt.Printf("%-30s%-15s%-15s%-5s%-5s\n", "Ime", "RD", "Datum", "Let", "Dni")
        crta(st)
        for _, k := range names {
            nm, rd, dt, yrs, dys := getData(k)
            fmt.Printf("%-30s%-15s%-15s%-5d%-5d\n", nm, rd, dt, yrs, dys)
        }
        crta(st)
    } else {
        fmt.Printf("\n" + "%s not in the list!\n", mynm)
        DisplayTen()
    }
}

func DisplayAll() {
    crta(st)
    fmt.Printf("%-30s%-15s%-15s%-5s%-5s\n", "Ime", "RD", "Datum", "Let", "Dni")
    crta(st)
    for k, _ := range seznam {
        nm, rd, dt, yrs, dys := getData(k)
        fmt.Printf("%-30s%-15s%-15s%-5d%-5d\n", nm, rd, dt, yrs, dys)
    }
    crta(st)
}

func DisplayTen() {
    crta(st)
    fmt.Printf("%-30s%-15s%-15s%-5s%-5s\n", "Ime", "RD", "Datum", "Let", "Dni")
    crta(st)
    i := 0
    for k, _ := range seznam {
        nm, rd, dt, yrs, dys := getData(k)
        fmt.Printf("%-30s%-15s%-15s%-5d%-5d\n", nm, rd, dt, yrs, dys)
        i += 1
        
        if i >= 9 {
            break
        }
    }
    crta(st)
}


func getData(mynm string) (string, string, string, int, int) {
    var myprs string
    var prsrd string
    var prsd string
    var prsyrs int
    var prsdni int
    
    
    for prsn, _ := range seznam {
        if strings.Contains(prsn, mynm) {
            prsrd = seznam[prsn]
            myprs = prsn
        }
    }
    
 
    thisY := time.Time.Format(danes, "2006")
    myRDs := seznam[myprs]
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

    return myprs, prsrd, prsd, prsyrs, prsdni
}


func main() {
    fjl := "ROJSTNIDNEVI.txt"
    file, err := os.Open(fjl)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    // populate map seznam with name, date entries
    for scanner.Scan() {
        krneki := strings.Split(scanner.Text(), ",")
        nm := krneki[0]
        dt := krneki[1]
        seznam[nm] = dt
    }

    // fmt.Println(seznam, "\n")
    Display("Mitja")
    
    Display("ŠŠpala")

    // DisplayAll()
}
