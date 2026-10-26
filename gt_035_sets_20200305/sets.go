// filename:	sets.go
// from: 		https://www.davidkaya.com/sets-in-golang/
//				file:///C:/Users/gregor.redelonghi/Dropbox/ODPRTO/_GO/go_sets_20191128.txt

package main

import "fmt"

// creating generic value for the map
// struct{} - type
// struct{}{} - empty struct --> size = 0!!
var exists = struct{}{}

// encapsulate map in its own struct
type set struct {
	m map[string]struct{}
}

func main() {
	s := NewSet()

	s.Add("Peter")
	s.Add("David")
	s.Add("Paul")
	s.Add("Matheus")
	s.Add2("Colin")
	s.Add2("James")
	s.Add2("David")
	s.Add2("Colin")
	s.Add2("Martin")

	fmt.Printf("Set contains Peter? ... %v\n", s.Contains("Peter"))
	fmt.Printf("Set contains Gegorge? ... %v\n", s.Contains("Peter"))
	fmt.Printf("Set contains Colin? ... %v\n", s.Contains("Colin"))
	fmt.Printf("Set contains David? ... %v\n", s.Contains("David"))
	fmt.Printf("Set contains Harrald? ... %v\n", s.Contains("Harrald"))

	s.Remove("David")
	fmt.Printf("Set contains Colin? ... %v\n", s.Contains("Colin"))
	fmt.Printf("Set contains David? ... %v\n", s.Contains("David"))

	s.Remove("David")
	fmt.Printf("Set contains David? ... %v\n", s.Contains("David"))

} // end main

func NewSet() *set {
	s := &set{}
	s.m = make(map[string]struct{})
	return s
}

func (s *set) Add(value string) {
	// instead of: s.m[value] = struct{}{}
	s.m[value] = exists
}

func (s *set) Add2(value string) {
	fmt.Printf("Adding \"%s\" to set ... ", value)

	if s.Contains(value) {
		fmt.Printf("\"%s\" already in set. Skipping.\n", value)
		return
	}
	s.m[value] = exists
	fmt.Print(" [OK]\n")
}

func (s *set) Remove(value string) {
	fmt.Printf("Removing \"%s\" from set ...", value)
	if not := s.Contains(value); not == false {
		fmt.Printf("\"%s\" NOT in set. Skipping.\n", value)
		return
	}
	delete(s.m, value)
	fmt.Print(" [OK]\n")
}

func (s *set) Contains(value string) bool {
	_, c := s.m[value]
	return c
}
