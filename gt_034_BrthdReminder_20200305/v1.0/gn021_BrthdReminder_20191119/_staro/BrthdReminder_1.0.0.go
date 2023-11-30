package main

// v1.0.0	(20191108/1/d -- ORIGINAL)
// v1.0.1   (

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var seznam = make(map[string]string)

func Display(mynm string) {
	for name, _ := range seznam {
		if strings.Contains(name, mynm) {
			prsData := seznam[name]
			fmt.Printf("RD for %s: %s\n", name, prsData)
			}
	}
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
    fmt.Println(seznam, "\n")
	Display("Špela")
	Display("Mitja")
}
