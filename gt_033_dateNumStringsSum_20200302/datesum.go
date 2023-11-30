package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	var dates = []string{"06000125", "19680222", "19991226", "20190131", "20200222", "20200223", "20200322", "20200323", "20200324"}
	var dates2 = []string{"25.01.0600", "22.02.1968", "26.12.1999", "31.01.2019", "22.02.2020", "23.02.2020", "22.03.2020", "23.03.2020", "24.03.2020"}

	for _, el := range dates {
		sum := sestej(el)
		fmt.Printf("Date: %8s --> sum: %9d\n", el, sum)
	}

	fmt.Println()

	for _, el2 := range dates2 {
		sum2 := sestej2(el2)
		fmt.Printf("Date2: %8s --> sum: %9d\n", el2, sum2)
	}

	var bigNum float64
	bigNum = math.Pow(2, 600)
	fmt.Println("\n" + "2**600:")
	fmt.Printf("%.0f\n", bigNum)
	fmt.Printf("%g\n", bigNum)

	// cmp bool := sum1 - sum2

} // end main

func sestej(date string) int {
	sum := 0
	yr_s := string(date[:4])
	mn_s := string(date[4:6])
	dy_s := string(date[6:])

	yr, _ := strconv.Atoi(yr_s)
	mn, _ := strconv.Atoi(mn_s)
	dy, _ := strconv.Atoi(dy_s)

	sum = (yr * 10000) + (mn * 100) + dy
	return sum
}

func sestej2(dateString string) int {
	sum2 := 0
	var date2 []string
	date2 = strings.Split(dateString, ".")

	yr_s := date2[2]
	mn_s := date2[1]
	dy_s := date2[0]

	yr, _ := strconv.Atoi(yr_s)
	mn, _ := strconv.Atoi(mn_s)
	dy, _ := strconv.Atoi(dy_s)

	sum2 = (yr * 10000) + (mn * 100) + dy
	return sum2
}
