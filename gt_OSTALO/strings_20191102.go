package main

import (
    "fmt"
    "strings"
)

func main() {

    fmt.Println("Printing multiple ocurencies of a character or string:")
    fmt.Println(strings.Repeat("s", 10))

    dm := "-#- "
    fmt.Println(strings.Repeat(dm, 10))

    S := strings.Repeat("*", 30)
    fmt.Println(S)

    Line := "My name is Gregor Redelonghi, and I come from Ljubljana, Slovenija. Welcome!"
    fmt.Printf("Original line: %s\n", Line)

    sv := "My"
    fmt.Printf("Original line starts with \"%s\"? ", sv)
    fmt.Printf("%t\n", strings.HasPrefix(Line, sv))

    sv = "Oh My"
    fmt.Printf("Original line starts with \"%s\"? ", sv)
    fmt.Printf("%t\n", strings.HasPrefix(Line, sv))

    sv = "Welcome!"
    fmt.Printf("Original line ends with \"%s\"? ", sv)
    fmt.Printf("%t\n", strings.HasSuffix(Line, sv))

    sv = "Ljubljana"
    fmt.Printf("Original line contains \"%s\"? ", sv)
    fmt.Printf("%t\n", strings.Contains(Line, sv))

    sv = strings.ToUpper("Ljubljana")
    fmt.Printf("Original line contains \"%s\"? ", sv)
    fmt.Printf("%t\n", strings.Contains(Line, sv))

    sv = strings.ToUpper("slovenija")
    fmt.Printf("Original line contains \"%s\" (no Case important)? ", sv)
    fmt.Printf("%t\n", strings.Contains(strings.ToUpper(Line), sv))

    fmt.Printf("Splitting original line on spaces: ")
    sv2 := strings.Split(Line, " ")
    fmt.Printf("[" + strings.Join(sv2, " | ") + "]")

}
