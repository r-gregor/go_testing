package main

import "fmt"

var myname string = "Gregor"
var my_sname string = "Redelonghi"

func numsList(border int) {
    for i := 0; i < border + 1; i++ {
        fmt.Printf("Element %d\n", i)
    }
}

func test(name string) {
	mjn := "Greg"
	if name == mjn {
		fmt.Println("OK!")
	} else {
		fmt.Printf("The name is not: %s!\n", mjn)
	}
}

func main() {
    fmt.Printf("My name is: %s, and my surname is: %s.\n", myname, my_sname)
    
    numsList(15)

	test("Zala")
	test("Greg")


}


