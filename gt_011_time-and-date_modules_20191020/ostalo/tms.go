// tms.go func(
package main

import (
  "fmt"
  "time"
)

func Day() string {
  now := time.Now()
  return now.Format("20060102")
}

func DayAndTime() string {
  now := time.Now()
  return now.Format("20060102_030405")
}

func tms() string {
  now := time.Now()
  timeAndDate := now.Format("20060102_030405")
  return fmt.Sprintf("[ %s ] --", timeAndDate)
}

func main() {
  Day := Day()
  fmt.Println("Current date:", Day)

  DayAndTime:= DayAndTime()
  fmt.Println("Current date and time:", DayAndTime)

  fmt.Printf("%s Starting ...\n", tms())

  for i := 1; i < 11; i++ {
    fmt.Printf("%s Iteration: %d\n", tms(), i)
    time.Sleep(1 * time.Second)
  }
}
