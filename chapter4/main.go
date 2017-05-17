package main

import (
    "fmt"
)

var global_x = "I am everywhere in this file"

func main() {
    // 4. Variables
    var x string = "Hello World"
    fmt.Println(x)

    var x_1 string
    x_1 = "Hello World"
    fmt.Println(x_1)

    var x_first = "first "
    fmt.Println(x_first)
    var x_second = "second"
    x_first += x_second
    fmt.Println(x_first)

    var xx string = "hello"
    var yy string = "world"
    fmt.Println(xx == yy)

    //var x_str = "String Hello World"
    x_num := 5
    fmt.Println(x_num)

    // 4.1 Naming your variable
    doggie_name := "Max"
    fmt.Println("My dog's name is ", doggie_name)
    
    // 4.2 Variable scoping
    fmt.Println(global_x)
    f()

    // 4.3 Constants
    const const_string = "I am 99.99 the same value as before"
    //const_string = "Yet another string, but will it work?"
    fmt.Println(const_string)

    // 4.4 Defining muliple variables
    // can use "var" or "const" for initializing multiple variables
    //var (
        //g = 5
        //h = 10
        //i = 15
    //)

    // 4.5 An example program
    fmt.Println("\n\n")
    fmt.Print("Enter a number: ")
    var input float64
    fmt.Scanf("%f", &input)

    output := input * 2

    fmt.Println(output)

    // Exercises:
    // -------------------- Exercise #1
    FarToCel()

    // ------------------- Enxercise #2
    FeetToMeters()
}

func f() {
    fmt.Println(global_x)
}

func FarToCel() {
    fmt.Println("Enter a number to be converted to Celsius: ")
    var fahrenheit float64
    fmt.Scanf("%f", &fahrenheit)

    result := (fahrenheit - 32) * 5/9

    fmt.Printf("Result is: %0.2f\n", result)
}

func FeetToMeters() {
    fmt.Println("Enter a number(in feet) to be converted to Meters: ")
    var feet float64
    metre := 0.3048
    fmt.Scanf("%f", &feet)

    result := feet * metre

    fmt.Printf("Result is: %0.4f\n", result)
}
