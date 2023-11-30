package main

import (
    "fmt"
    "time"
)

func Action() {
    fmt.Println("Action going on for 3 seconds")
    time.Sleep(3 * time.Second)
}

func main() {
    
    start := time.Now()
    fmt.Printf("start time: %s\n", start.Format("15:04:05.999"))
    
    // run action
    Action()
    
    end := time.Now()
    fmt.Printf("end time: %s\n", end.Format("15:04:05.999"))
    
    diff := end.Sub(start)
    fmt.Printf("time difference: %s\n", diff)

}