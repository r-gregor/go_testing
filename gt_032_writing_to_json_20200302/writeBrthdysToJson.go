// filename: writeBrthdysToJson.go

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Person struct {
	Name  string `json: "name"`
	Brthd string `json: "birthday"`
}

type Persons struct {
	List []Person `json: "list"`
}

func main() {

	var persons Persons
	persons.List = append(persons.List, Person{"Gregor Redelonghi", "22.02.1968"})
	persons.List = append(persons.List, Person{"Zala Redelonghi", "07.05.1996"})
	persons.List = append(persons.List, Person{"Mark Redelonghi", "17.04.1999"})
	persons.List = append(persons.List, Person{"Tadeja Mali Redelonghi", "19.01.1970"})
	persons.List = append(persons.List, Person{"Špela Redelonghi", "11.02.2008"})

	/* TEST
	fmt.Println(persons)
	persons.DisplayAll()
	*/

	fmt.Print("Exporting data to brthdys.json ... ")
	// data, _ := json.Marshal(persons) // not indented
	data, _ := json.MarshalIndent(persons, "", " ")

	/* TEST
	fmt.Println(string(data))
	*/

	var jsonFile string = "brthdys.json"
	checkFile(jsonFile)

	err2 := ioutil.WriteFile(jsonFile, data, 0644)

	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println("Write successful!")
	}

	persons.DisplayAll()

} // end main

func (l Persons) DisplayAll() {
	for _, prsn := range l.List {
		fmt.Printf("\nName: %s\nBirthday: %s\n", prsn.Name, prsn.Brthd)
	}
}

func checkFile(filename string) {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		fmt.Println("\n"+"There is NO file:", filename)
		os.Exit(1)
	}
	return
}

// full checkFile version -- return error !!
/*
func checkFile(filename string) error {
    _, err := os.Stat(filename)
    if os.IsNotExist(err) {
        _, err := os.Create(filename)
        if err != nil {
            return err
        }
    }
    return nil
}
*/
