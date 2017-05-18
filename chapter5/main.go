package main

import (
    "fmt"
)

func main() {
    i := 1
    for i <= 10 {
        fmt.Println(i)
        i += 1
    }

    val_is_now()
    multiple_init()
    break_key()
    continue_key()
    range_key()
}

func val_is_now() {
    for i := 0; i < 5; i++ {
        fmt.Println("Value of i is now: ", i)
    }
}

func multiple_init() {
    // multiple initialization; a consolidated bool expression with && and ||
    // multiple incrementation

    var (
        i = 0
        j = 5
        s = "a"
    )

    for i < 3 && j < 100 && s != "aaaaa" {
        fmt.Println("Value of i, j, s:", i, j, s)
        i, j, s = i+1, j+1, s + "a"
    }
}

func break_key() {
    i := 0
    for { // since there are no checks, this is an infinite loop
        if i >= 3 { break }
        fmt.Println("Value of i is: ", i)
        i++;
    }
    fmt.Println("A statement just after for loop.")
}

func continue_key() {
    // a continue within this loop will bring back execution to the beginning
    // of the loop. Check and increments in for loop will be executed.
    for i := 0; i < 7; i++ {
        // control comes here when there is a 'continue' within this for block
        if i % 2 == 0 {
            continue // if it is an even number, go back to beginning of for loop
        }
        fmt.Println("Odd: ", i) //execution will reach here only when i%2 is not 0, and therefore it is odd
    }
}

func range_key() {
    //on an array, range returns the index
    a := [...]string{"a", "b", "c", "d"}
    for i := range a {
        fmt.Println("Array item", i, "is", a[i])
    }

    // on a map, range return the key
    capitals := map[string] string {"France":"Paris", "Italy":"Rome", "Japan":"Tokyo"}
    for key := range capitals {
        fmt.Println("Map item: Capital of", key, "is", capitals[key])
    }

    // range can also return two items, the index/key and the corresponding value
    for key2, val := range capitals {
        fmt.Println("Map item: Capital of", key2, "is", val)
    }
}
