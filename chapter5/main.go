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
    if_key()
    if_key_iteration()
    switch_key()
    big_or_small()
    evenly_divisible_by_3()
    FizzBuzz()
    iota_test()
    test_memory_address()
    test_pointers()
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

func if_key() {
    // odd or even using if
    i := 10;

    if i % 2 == 0 {
        fmt.Println("Is divisible by 2");
    } else if i % 3 == 0 {
        fmt.Println("Is divisible by 3");
    } else if i % 4 == 0 {
        fmt.Println("Is divisible by 4")
    }
}

func if_key_iteration() {
    for i := 1; i <= 100; i++ {
        if i % 2 == 0 {
            fmt.Println(i, "even");
        } else {
            fmt.Println(i, "odd");
        }
    }
}

func switch_key() {

    i := 4;

    switch i {
        case 0: fmt.Println("Zero:");
        case 1: fmt.Println("One:");
        case 2: fmt.Println("Two:");
        case 3: fmt.Println("Three");
        case 4: fmt.Println("Four");
        case 5: fmt.Println("Five");
        default: fmt.Println("Unknown Number");
    }
}


// page 57 golang-book exercises
// #1
func big_or_small() {
    i := 10

    if i > 10 {
        fmt.Println("Big");
    } else {
        fmt.Println("Small");
    }
}

/* 
    Write a program thhat prints out all the numbers evenly divisible by 3 between 1 and 100. numbers such as (3, 6, 9, ... etc.)
*/
func evenly_divisible_by_3() {
    for i := 1; i <= 100; i++ {
        if i % 3 == 0 {
            fmt.Println(i)
        }
    }
}

/*
    Write a program that prints the number from 1 to 100. But for multiples of
    three print "Fizz" instead of the number and for the multiples of five print
    "Buzz". For numbers which are multiples of both three and five print
    "FizzBuzz".
*/
func FizzBuzz() {

    fmt.Println("FizzBuzz() Program starts here -------------------- ");
    for i := 1; i <= 100; i++ {
        if i % 15 == 0 {
            fmt.Println(i, "FizzBuzz");
        } else if i % 3 == 0 {
            fmt.Println(i, "Fizz");
        } else if i % 5 == 0 {
            fmt.Println(i, "Buzz")
        } else {
            fmt.Println(i, "n/a")
        }
    }
}

func iota_test() {

    const (
        Sunday      = iota
        Monday
        Tuesday
        Wednesday
        Thursday
        ThankGodItsFriday
        PartyDay
        numberOfDays
    )

    fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, ThankGodItsFriday, PartyDay, numberOfDays)

    type ByteSize float64

    const (
        _           = iota // ignore first line by assigning to blank identifier
        KB ByteSize = 1 << (10 * iota)
        MB
        GB
        TB
        PB
        EB
        ZB
        YB
    )

    fmt.Println(KB, MB, GB, TB, PB, EB, ZB, YB)

    const (
        c0  = 1 << iota
        c1  = 1 << iota
        c2  = 1 << iota
        c3  = 1 << iota
    )

    const (
        u   = iota * 42
        v float64 = iota * 42
        w   = iota * 42
    )

    fmt.Println(c0, c1, c2, c3);
    fmt.Println(u, v, w);
}

func test_memory_address() {
    a := 43

    fmt.Println("a - ", a);
    fmt.Println("a's memory resides @ ", &a);

    const metersToYards float64 = 1.09361 // should have been globally instantiated

    var meters float64
    fmt.Println("Enter meters swam: ")
    fmt.Scan(&meters)
    yards := meters * metersToYards
    fmt.Println(meters, "meters is", yards, "yards.");
}

func test_pointers() {

    a := 43;

    fmt.Println(a);
    fmt.Println(&a);

    var b *int = &a; // b is referencing it
    fmt.Println(b);
    fmt.Println(*b); // dereference

    *b = 100;
    fmt.Println(a);

    // var c int = &a;
}
