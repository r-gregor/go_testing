// tmsmain.go
package main

import (
  "fmt"
  "time"
  tp "github.com/r-gregor/tmspkg"
)

func main() {

  fmt.Println("Current date:", tp.Day())

  fmt.Println("Current date and time:", tp.DayAndTime())

  fmt.Printf("%s Starting ...\n", tp.Tms())

  for i := 1; i < 11; i++ {
    fmt.Printf("%s Iteration: %d\n", tp.Tms(), i)
    time.Sleep(1 * time.Second)
  }
}
