package main

import (
    "fmt"
)

var global_x = "Hello World"

func main() {
    // 1. Numbers  section
    fmt.Println("1 + 1 = ", 1.0 + 1.0)

    // 2. Strings section
    strVal := "This is a Hello World"
    fmt.Println(len(strVal))
    fmt.Println("ßello World"[0])
    fmt.Println("Hello " + "World")

    // 3. Booleans section
    fmt.Println(true && true)
    fmt.Println(true && false)
    fmt.Println(true || true)
    fmt.Println(true || false)
    fmt.Println(!true)
    fmt.Println(false || false || true)
}
