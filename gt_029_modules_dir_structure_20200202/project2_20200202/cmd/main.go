package main

import (
	"fmt"

	"github.com/r-gregor/project2_20200202/pkg/mymodule"
	sb "github.com/r-gregor/project2_20200202/pkg/mymodule/subdir1/subsubdir1"
	tm "github.com/r-gregor/project2_20200202/pkg/tmspkg"
)

func main() {
	fmt.Println("This is a test executable using tmspkg")
	fmt.Printf("Today: %s\n", tm.Day())
	fmt.Printf("DayAndTime: %s\n", tm.DayAndTime())
	fmt.Printf("Tms: %s\n", tm.Tms())

	mymodule.PrintGreeting("Gregor Redelonghi")

	sb.HeyYou()

}
