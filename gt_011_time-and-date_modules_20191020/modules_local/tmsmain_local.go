// tmsmain_local.go
package main

import (
  "fmt"
  "time"
  "github.com/r-gregor/tmspkg"
)

func main() {

  fmt.Println("Current date:", tmspkg.Day())

  fmt.Println("Current date and time:", tmspkg.DayAndTime())

  fmt.Printf("%s Starting ...\n", tmspkg.Tms())

  for i := 1; i < 11; i++ {
    fmt.Printf("%s Iteration: %d\n", tmspkg.Tms(), i)
    time.Sleep(1 * time.Second)
  }
}
