package main

import (
    hellomod "github.com/donvito/hellomod"
    hellomodV2 "github.com/donvito/hellomod/v2"
)

func main() {
    hellomod.SayHello()
    hellomodV2.SayHello("Melvin")
}
