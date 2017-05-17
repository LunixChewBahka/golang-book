package main

import "fmt"

func main() {
    fmt.Println("1 + 1 = ", 1.0 + 1.0)

    // Strings section
    fmt.Println(len("Hello World"))
    fmt.Println("ßello World"[0])
    fmt.Println("Hello " + "World")

    // Booleans section
    fmt.Println(true && true)
    fmt.Println(true && false)
    fmt.Println(true || true)
    fmt.Println(true || false)
    fmt.Println(!true)
}
