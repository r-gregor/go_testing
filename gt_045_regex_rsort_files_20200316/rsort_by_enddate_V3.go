// fname: rsort_by_enddate_V2.go
// updated 20211020

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

var fileNamesByDate = make(map[string][]string)

func main() {

	var path string

	if len(os.Args) == 2 {
		path = os.Args[1]
	} else {
		// get current working directory path
		path1, err1 := os.Getwd()
		if err1 != nil {
			log.Fatal(err1)
		}
		path = path1
	}

	//
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

			myFile := fnameinfo{Fdate: date, Fname: fname}
			fileNamesByDate[date] = append(fileNamesByDate[date], myFile.Fname)
		}

	}

	// a slice of keys
	keys := make([]string, 0, len(fileNamesByDate))
	for k := range fileNamesByDate {
		keys = append(keys, k)
	}

	// reverse sorted slice of keys
	sort.Sort(sort.Reverse(sort.StringSlice(keys)))

	// print workinkdir
	// fmt.Printf("Working dir: %s\n\n", path)

	// print sorted by date
	for _, k := range keys {
		for _, el := range fileNamesByDate[k] {
			YR := k[0:4]
			MN := k[4:6]
			DY := k[6:]
			fmt.Printf("%s-%s-%s %s\n", YR, MN, DY, el)
		}
	}
} // end main
