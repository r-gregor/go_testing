// tmspkg.go
package tmspkg

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

func Tms() string {
  now := time.Now()
  timeAndDate := now.Format("20060102_030405")
  return fmt.Sprintf("[ %s ] --", timeAndDate)
}
