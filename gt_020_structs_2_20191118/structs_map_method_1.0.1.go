package main

// v1.0.0	(20191105/1/d)	-- ORIGINAL
// v1.0.1	(20191105/2/d)	-- cosmetics ...

import "fmt"

type data struct {
	m map[int]string
}

func main() {
	// without method MakeMap()
	// d1 := data{m: make(map[int]string)}

	d1 := new(data).MakeMap()

	d1.DisplayDataInMap()

	d1.AddToMap(1, "KAKA")
	d1.AddToMap(1, "KAKO")
	d1.AddToMap(2, "KAKO")
	d1.AddToMap(3, "TINY")
	d1.AddToMap(4, "MINY")
	d1.AddToMap(5, "MOE")

	fmt.Println(d1.FindInMap(15))
	fmt.Println(d1.FindInMap(5))

	d1.DisplayDataInMap()

}

func (*data) MakeMap() data {
	return data{m: make(map[int]string)}
}

func (d *data) AddToMap(i int, s string) {
	if d.m[i] != "" {
		fmt.Printf("Trying to insert \"%s\" into key %d ... ", s, i)
		fmt.Printf("Key %d already taken\n", i)
	} else {
		fmt.Printf("Adding %d ==> \"%s\"\n", i, s)
		d.m[i] = s
	}
}

func (d *data) FindInMap(i int) string {
	result := ""
	if d.m[i] == "" {
		result = fmt.Sprintf("Key %d NOT in map, or empty.\n", i)
	} else {
		result = fmt.Sprintf("Value for key %d is: \"%s\"\n", i, d.m[i])
	}
	return result
}

func (d *data) DisplayDataInMap() {
	if len(d.m) == 0 {
		fmt.Println("A map is empty!")
	} else {
		fmt.Println("Displaying all elemnts of map ...")
		for k, v := range d.m {
			fmt.Printf("Key: %d\t value: %s\n", k, v)
		}
	}
}
