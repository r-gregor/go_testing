// eval.go

package main

import (
	"fmt"
)

func main() {

	var num string = "15"
	var num2 string = "25"

	var line string = "System upload at: %-" + num + "shours.\n"
	sentence1 := fmt.Sprintf(line, "14:00:00")

	fmt.Println("line1:")
	fmt.Print(sentence1)
	showChars(sentence1)
	fmt.Printf("The length oh line1 is %d.\n", len(sentence1))

	// fmt.Println(strings.Repeat("-", 80))
	fmt.Println()
	fmt.Println()

	line2 := "System upload at: %-" + num2 + "shours.\n"
	sentence2 := fmt.Sprintf(line2, "23:59:59")

	fmt.Println("line2:")
	fmt.Print(sentence2)
	showChars(sentence2)
	fmt.Printf("The length oh line2 is %d.\n", len(sentence2))

}

func showChars(line string) {

	for i := 0; i < len(line)-1; i++ {
		if string(line[i]) == " " {
			fmt.Print("-")
		} else {
			fmt.Print("^")
		}
	}
	fmt.Println()
}
