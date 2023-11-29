package main

// v1.0.0	20200406/01/d
// v1.0.1	20200407/02/en	-- absolute path to country-capitals.json
//                             can be called from anywhere !! -- NOT WORKING if build directory changed!!
//							-- countriesMap --> map[string]Country

/*
 * TODO:
 * put countries data from Countries.CList into map: map[string]Country:
 * - key:	string: country name
 * - value:	struct Country to iterate over data ...
 *
 */

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
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

func main() {

	countriesMap := make(map[string]Country)

	// v1.0.1
	bin, _ := os.Executable()
	pth, _ := filepath.Abs(filepath.Dir(bin))
	// test
	// fmt.Println(pth)

	// v1.0.8
	fjl := filepath.Join(pth, "country-capitals.json")
	// test
	// fmt.Println(fjl)

	// v1.0.1
	// jsonFile, err := os.Open("country-capitals.json")
	jsonFile, err := os.Open(fjl)

	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println("Successfully Opened users.json")
	// defer the closing of jsonFile so that it can be parsed later on
	defer jsonFile.Close()

	// read file as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// initialize the Countries struct
	var countries Countries

	// json to Countries.CList magic
	json.Unmarshal([]byte(byteValue), &countries.CList)

	// v1.0.0
	// test
	// fmt.Println(countries)
	// countries.DisplayAll()

	// v1.0.1
	countries.loadToCountriesMap(countriesMap)

	// TEST
	// fmt.Println(countriesMap)

	// RUN
	displayAllCountriesMap(countriesMap)
} // end main

func (c Countries) DisplayAll() {
	for i := 0; i < len(c.CList); i++ {
		fmt.Println("Country: " + c.CList[i].CountryName)
		fmt.Println("Capital: " + c.CList[i].CapitalName)
		fmt.Println("Country code: " + c.CList[i].CountryCode)
		fmt.Println("Continent: " + c.CList[i].ContinentName)
		fmt.Println("---------------------------------------")
	} // end for
} // end func

func (c Countries) loadToCountriesMap(cmap map[string]Country) {
	for i := 0; i < len(c.CList); i++ {
		cName := c.CList[i].CountryName

		cmap[cName] = c.CList[i]
	} // end for
} // end func

func displayAllCountriesMap(cmap map[string]Country) {
	for _, cntr := range cmap {
		fmt.Printf("Country: %s\n", cntr.CountryName)
		fmt.Printf("Capital: %s\n", cntr.CapitalName)
		fmt.Printf("Country code: %s\n", cntr.CountryCode)
		fmt.Printf("Continent: %s\n", cntr.ContinentName)
		fmt.Println("---------------------------------------")
	} // end for
} // end func
