package main

import (
	"fmt"
	"github.com/r-gregor/projectName3/modules/pkg1"
	"github.com/r-gregor/projectName3/modules/pkg2"
)

func main() {
	pkg1.Info1()
	pkg2.Info2()

	x := 3
	y := 2

	fmt.Printf("The sum of %d and %d is: %d\n", x, y, pkg1.Addition(x, y))
}
