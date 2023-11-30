// v1.0.0	(20191117/1/d	-- ORIGINAL)

package main

import (
	"fmt"
	"strings"
)

type judoka struct {
	name   string
	belt   string
	rank   int
	hobbys []string
}

func main() {

	g := judoka{
		name:   "Gregor",
		belt:   "black",
		rank:   100,
		hobbys: []string{"judo", "Hiking", "Dancing"},
	}

	g.PrintInfo()

	g.ChangeRank(50)
	g.PrintInfo()

	g.ChangeRank2(50)
	g.PrintInfo()

	g.AddHobby("Swimming")
	g.PrintInfo()

}

func (jk judoka) PrintInfo() {
	fmt.Printf("Judoka:\t%s\nbelt:\t%s\nrank:\t%d\n", jk.name, jk.belt, jk.rank)
	hbs := strings.Join(jk.hobbys, ", ")
	fmt.Printf("Hobys:\t%s\n", hbs)
	fmt.Println("----------------------------------")
}

func (jk judoka) ChangeRank(newrank int) {
	jk.rank = newrank
	fmt.Printf("Cannging rank of %s to: %d\n", jk.name, jk.rank)
}

func (jk *judoka) ChangeRank2(newrank int) {
	jk.rank = newrank
	fmt.Printf("Cannging rank of %s to: %d\n", jk.name, jk.rank)
}

func (jk *judoka) AddHobby(newhb string) {
	jk.hobbys = append(jk.hobbys, newhb)
	fmt.Printf("Adding %s to: %s's hobbys\n", newhb, jk.name)
}
