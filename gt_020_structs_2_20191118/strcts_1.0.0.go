package main

// v1.0.0	(20191117/1/d	-- ORIGINAL)

import "fmt"

type bboy struct {
	name  string
	age   int
	hobby []string
}

func main() {
	greg := bboy{name: "Gregor", age: 52}
	fmt.Println(greg)

	greg.ChangeName("RogerG")
	fmt.Println(greg)

	greg.hobby = append(greg.hobby, "judo")
	greg.hobby = append(greg.hobby, "dancing")
	fmt.Println(greg)

	greg.UpdateHobby("Hiking")
	fmt.Println(greg)

	greg.ChangeHobby(1, "CODING")
	fmt.Println(greg)

	greg.name = "Roudi"
	fmt.Println(greg)

	fmt.Println("123456789000")

	ChangeName2(greg)
	fmt.Println(greg)

	ChangeName3(&greg)
	fmt.Println(greg)

}

// ChangeName	change the name of pointer to bboy struct
func (b *bboy) ChangeName(nname string) {
	b.name = nname
}

// ChangeName2 change the name of name field of bboy struct directly (no pointer)
func ChangeName2(b bboy) {
	b.name = "KuKU"
}

// ChangeName3 change the name of pointer to bboy struct
func ChangeName3(b *bboy) {
	b.name = "KuKU"
}

// UpdateHobby update/add hobby to hobby field using pointer to struct
func (b *bboy) UpdateHobby(nhobby string) {
	b.hobby = append(b.hobby, nhobby)
}

// UpdateHobby CHANGE hobby field using pointer to struct
func (b *bboy) ChangeHobby(index int, nhobby string) {
	b.hobby[index] = nhobby
}
