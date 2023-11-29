package main

// v1.0.0   20200406/01/d
// v1.0.1   20200407/02/en   -- absolute path to country-capitals.json
//                              can be called from anywhere !! -- NOT WORKING if build directory changed!!
//                           -- countriesMap --> map[string]Country
//
// v1.0.2   20200407/03/en    - countriesMap as type
// v1.0.3   20200407/04/en    - necessery start up commands in initialize() function:
//                              initialize() returns countriesMap (type CountriesMap map[string]Country) needed in:
//                                     countriesMap.displayAllCountries() and countriesMap.displaySingle(patt)
//

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Countries struct {
	CList []Country
}

type Country struct {
	CountryName   string `json: CountryName`
	CapitalName   string `json: CapitalName`
	CountryCode   string `json: CountryCode`
	ContinentName string `json: ContinentName`
}

type CountriesMap map[string]Country

func main() {

	// v1.0.2
	var usage string = `
+-------------------------------------------------------------+
| Usage:                                                      |
| $> CountryCapitals.exe [ country name part ] or [ full ]    |
|       country name ... displays countries with part in name |
|       full         ... displays ALL countries (full list)   |
+-------------------------------------------------------------+
`
	// needed in functions countriesMap.displayAllCountries() and countriesMap.displaySingle(patt)
	countriesMap := initialize()

	// v1.0.2
	if len(os.Args) != 2 {
		fmt.Println(usage)
		os.Exit(1)
	} else {
		var patt string = os.Args[1]
		if strings.ToUpper(patt) == "FULL" {
			countriesMap.displayAllCountries()
		} else {
			countriesMap.displaySingle(patt)
		}
	}

} // end main

func (cmap CountriesMap) loadToCountriesMap(c Countries) {
	for i := 0; i < len(c.CList); i++ {
		cName := c.CList[i].CountryName

		cmap[cName] = c.CList[i]
	} // end for
} // end func

func (cntr CountriesMap) displayAllCountries() {
	for _, cntr := range cntr {
		fmt.Printf("Country: %s\n", cntr.CountryName)
		fmt.Printf("Capital: %s\n", cntr.CapitalName)
		fmt.Printf("Country code: %s\n", cntr.CountryCode)
		fmt.Printf("Continent: %s\n", cntr.ContinentName)
		fmt.Println("---------------------------------------")
	} // end for
} // end func

func (cntr CountriesMap) displaySingle(patt string) {
	var count int = 0
	for _, cntr := range cntr {
		var mypatt = strings.ToUpper(patt)
		var currentC = strings.ToUpper(cntr.CountryName)

		if strings.Contains(currentC, mypatt) {
			count += 1
			fmt.Printf("Country: %s\n", cntr.CountryName)
			fmt.Printf("Capital: %s\n", cntr.CapitalName)
			fmt.Printf("Country code: %s\n", cntr.CountryCode)
			fmt.Printf("Continent: %s\n", cntr.ContinentName)
			fmt.Println("---------------------------------------")
		} else {
			continue
		} // end if

	} // end for

	if count == 0 {
		fmt.Printf("No country name that contains [%s] found!\n", patt)
		os.Exit(1)
	}

} // end func

func initialize() CountriesMap {
	var countries Countries
	var cmap = make(CountriesMap)

	bin, _ := os.Executable()
	pth, _ := filepath.Abs(filepath.Dir(bin))

	fjl := filepath.Join(pth, "country-capitals.json")

	jsonFile, err := os.Open(fjl)

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal([]byte(byteValue), &countries.CList)

	cmap.loadToCountriesMap(countries)

	return cmap
}
