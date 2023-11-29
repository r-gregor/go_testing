package main

// v1.0.0	20200406/01/d	-- ORIGINAL

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
)

type Countries struct {
	// countries []Country
	CList []Country
}

type Country struct {
	CountryName   string `json: CountryName`
	CapitalName   string `json: CapitalName`
	CountryCode   string `json: CountryCode`
	ContinentName string `json: ContinentName`
}

func main() {

	// Open our jsonFile
	jsonFile, err := os.Open("country-capitals.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read file as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var countries Countries

	json.Unmarshal([]byte(byteValue), &countries.CList)

	// test
	// fmt.Println(countries)
	countries.DisplayAll()

} // end main

func (c Countries) DisplayAll() {
	for i := 0; i < len(c.CList); i++ {
		fmt.Println("Country: " + c.CList[i].CountryName)
		fmt.Println("Capital: " + c.CList[i].CapitalName)
		fmt.Println("Country code: " + c.CList[i].CountryCode)
		fmt.Println("Continent: " + c.CList[i].ContinentName)
		fmt.Println("---------------------------------------")
	}
}
