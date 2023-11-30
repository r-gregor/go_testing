package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	path, err1 := os.Getwd()
	if err1 != nil {
		log.Fatal(err1)
	}

	// test
	// fmt.Println(path)

	file, err2 := os.Open(path)
	if err2 != nil {
		log.Fatal(err2)
	}

	names, err3 := file.Readdirnames(0)
	if err3 != nil {
		log.Fatal(err3)
	}

	// test
	// fmt.Println(names)

	/*
		for _, fname := range names {
			if strings.Contains(fname, "txt") {
				fmt.Printf("Filename: %s\n", fname)
			}

		}
	*/

	for _, fname := range names {
		// matched, _ := regexp.MatchString("[[:digit:]]{8}", fname)
		// better: string litteral: `` -- \d NOT working when between "" so have to use [[:digit:]]
		matched, _ := regexp.MatchString(`\d{8}`, fname)
		if matched {
			re, _ := regexp.Compile(`\d{8}`)
			date := re.FindString(fname)
			fmt.Printf("%s %s\n", date, fname)
		}

	}

}
