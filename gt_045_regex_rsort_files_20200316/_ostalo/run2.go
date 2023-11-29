// fname: run2.go

package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
)

type fnameinfo struct {
	Fdate string
	Fname string
}

var datesSet []string
var listOfFnames []fnameinfo

func main() {

	path, err1 := os.Getwd()
	if err1 != nil {
		log.Fatal(err1)
	}

	file, err2 := os.Open(path)
	if err2 != nil {
		log.Fatal(err2)
	}

	names, err3 := file.Readdirnames(0)
	if err3 != nil {
		log.Fatal(err3)
	}

	for _, fname := range names {
		matched, _ := regexp.MatchString(`\d{8}`, fname)
		if matched {
			re, _ := regexp.Compile(`\d{8}`)
			date := re.FindString(fname)

			// test
			// fmt.Printf("%s %s\n", date, fname)

			myFile := fnameinfo{Fdate: date, Fname: fname}
			addFnameToList(myFile)
			addDateToSet(date)
		}

	}

	showListOfFnames()

	// sort.Strings(datesSet)
	sort.Sort(sort.Reverse(sort.StringSlice(datesSet)))
	showDatesSet()
} // end main

func addFnameToList(myfile fnameinfo) {
	listOfFnames = append(listOfFnames, myfile)
}

func showListOfFnames() {
	for _, fn := range listOfFnames {
		fmt.Printf("%s %s\n", fn.Fdate, fn.Fname)
	}
}


func addDateToSet(date string) {

	var exists bool = false
	for _, d := range datesSet {

		if date == d {
			exists = true
			break
		} else {
			exists = false
		}
	}

	if !exists {
		fmt.Printf("Adding %s .to datessSet\n", date)
		datesSet = append(datesSet, date)
	}
}


func showDatesSet() {
	for _, d := range datesSet {
		fmt.Printf("%s\n", d)
	}
}

/*
func showSorted() {
	for _, d in dateSet {
		for ...
	}
}
*/
