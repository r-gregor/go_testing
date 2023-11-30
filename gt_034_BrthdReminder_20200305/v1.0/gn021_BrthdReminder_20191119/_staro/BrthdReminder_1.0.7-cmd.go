package main

// v1.0.0       (20191108/1/d   -- ORIGINAL)
// v1.0.1       (20191108/2/en)
// v1.0.2       (20191108/3/en)
// v1.0.3       (20191108/4/en  -- structs)
// v1.0.4       (20191110/5/d   -- sort)
// v1.0.5       (20191111/6/en  -- print those with [howmany] days till BD)
//                              -- skupni format za fmt.Printf --> frmt := ...
//                              -- BARVE
//                              -- cli argument (name)
// v1.0.5-cmd   (20191111/7/en) -- removed BARVE for display on cmd
//
// v1.0.6-cmd  (20191112/7/en) -- repair date overflip for diff == 01
// v1.0.7   (20191112/8/en) -- absolute path to ROJSTNIDNEVI.txt
//                             can be called from anywhere !!

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"          // v1.0.7
	"path/filepath" // v1.0.7
	"runtime"       // v1.0.7
	"sort"          // v1.0.4
	"strconv"
	"strings"
	"time"
)

type person struct {
	nm  string
	rd  string
	dt  string
	yrs int
	dys int
}

var danes = time.Now()
var seznam = make(map[string]string)
var dni int

var seznam2 []person

var st int = 30 + 15 + 15 + 10 + 5 + 4

// v1.0.5
var frmt = "%-30s%-15s%-15s%-10v%-5s%-4v\n"

func crta(n int) {
	for i := 0; i < n; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}

// Sort by dys, keeping original order or equal elements.
// v1.0.4
func sortMe() {
	sort.SliceStable(seznam2, func(i, j int) bool {
		return seznam2[i].dys < seznam2[j].dys
	})
}

func fillData() {

	var prsd string
	var prsyrs int
	var prsdni int
	var fjl string

	// get absolute path of mainprogam and ROJSTNIDNEVI.txt file
	_, filename, _, _ := runtime.Caller(0)
	absp, err := filepath.Abs(path.Dir(filename))

	fjl = absp + "\\ROJSTNIDNEVI.txt"

	// fjl := "./ROJSTNIDNEVI.txt"
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

		thisDateList := strings.Split(time.Time.Format(danes, "02.01.2006"), ".")
		thisY := thisDateList[2]

		myRDs := item.rd
		myRD := strings.Split(myRDs, ".")
		myRDThisYs := fmt.Sprintf("%s.%s.%s", myRD[0], myRD[1], thisY)
		danes2 := danes.AddDate(1, 0, 0)

		var YY string
		if myRD[0] < thisDateList[0] || myRD[1] < thisDateList[1] {
			YY = time.Time.Format(danes2, "2006")
		} else {
			YY = time.Time.Format(danes, "2006")
		}

		myRDThisYs = fmt.Sprintf("%s.%s.%s", myRD[0], myRD[1], YY)

		prsd = myRDThisYs
		myRDfinal, _ := time.Parse("02.01.2006", myRDThisYs)

		diff := myRDfinal.Sub(danes)

		//v1.0.06
		if myRDThisYs == time.Time.Format(danes, "02.01.2006") {
			dni = 0
		} else {
			// v1.0.5 ( +1)
			dni = int(diff.Hours()/24) + 1
		}

		now, _ := strconv.Atoi(YY)
		then, _ := strconv.Atoi(myRD[2])

		prsyrs = now - then
		prsdni = dni

		item.dt = prsd
		item.yrs = prsyrs
		item.dys = prsdni

		// test
		seznam2 = append(seznam2, *item)
		sortMe()

	}
}

func color(days int) string {
	color := ""
	if days <= 31 && days > 7 {
		color = "*"
	} else if days <= 7 && days > 0 {
		color = "**"
	} else if days == 0 {
		color = "***"
	} else {
		color = ""
	}
	return color
}

func DisplayAll() {
	crta(st)
	fmt.Printf(frmt, "Ime", "RD", "Datum", "Let", "Dni do RD", "")
	crta(st)
	for _, item := range seznam2 {
		fmt.Printf(frmt, item.nm, item.rd, item.dt, item.yrs, color(item.dys), item.dys)
	}
	crta(st)
}

// v1.0.4
func DisplayTen() {
	crta(st)
	fmt.Printf(frmt, "Ime", "RD", "Datum", "Let", "Dni do RD", "")
	crta(st)
	i := 0
	for _, item := range seznam2 {
		fmt.Printf(frmt, item.nm, item.rd, item.dt, item.yrs, color(item.dys), item.dys)
		i += 1
		if i >= 10 {
			break
		}
	}
	crta(st)
	fmt.Println("Prikaz prvih 10 oseb z najmanj dni do RD\n\n")
}

// v1.0.5
func Display(howmany int) {
	crta(st)
	fmt.Printf(frmt, "Ime", "RD", "Datum", "Let", "Dni do RD", "")
	crta(st)
	for _, item := range seznam2 {
		if item.dys < howmany {
			fmt.Printf(frmt, item.nm, item.rd, item.dt, item.yrs, color(item.dys), item.dys)
		} else {
			continue
		}
	}
	crta(st)
	fmt.Printf("Prikaz oseb z manj kot %d dni do RD\n\n", howmany)
}

// v1.0.5
func DisplayPers(selected []person) {
	crta(st)
	fmt.Printf(frmt, "Ime", "RD", "Datum", "Let", "Dni do RD", "")
	crta(st)
	for _, item := range selected {
		fmt.Printf(frmt, item.nm, item.rd, item.dt, item.yrs, color(item.dys), item.dys)
	}

	crta(st)
	// fmt.Printf("Prikaz oseb z manj kot %d dni do RD\n\n")
}

// MAIN
func main() {

	// v1.0.5
	fmt.Printf("\nDANES: \t%s\n", time.Time.Format(danes, "02.01.2006"))

	// get data from external file
	fillData()

	var single []person

	// v1.0.5
	if len(os.Args) == 2 {
		name := os.Args[1]
		if strings.ToUpper(name) == "ALL" {
			DisplayAll()
		} else {
			for _, el := range seznam2 {
				if strings.Contains(strings.ToUpper(el.nm), strings.ToUpper(name)) {
					single = append(single, el)
				} else {
					continue
				}
			}

			if len(single) != 0 {
				DisplayPers(single)
				fmt.Printf("Prikaz oseb ki vsebujejo \"%s\"\n\n", name)
			} else {
				Display(50)
			}
		}
	} else {
		Display(100)
	}

}
