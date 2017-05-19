/*
Arrays, Slices and Maps
*/
package main

import (
    "fmt"
)

func main() {
    /* 6.1 Arrays Section
            An array is a numbered sequence of eleemnts of a single type
            with a fixed length. In go they look like this:
    */
    var x [5]int
    x[4] = 100 // set the 5th element of array x to 100
    fmt.Println(x) // index starts at "0" where "0" is index[1] ; [0 0 0 0 100]
    fmt.Println(x[4]) // it only outputs the actual value not the array ; 100

    example_array1()
    example_array2()

    /* 6.2 Slices Section
        A slice is a segment of an array. Like arrays slices are indexable and
        have a length. Unlike arrays this length is allowed to change. Here's
        an example of a slice:
    */
    //var y []float64 // y has been created with a length of 0
    // should use the built-in function to create a slice
    //z := make([]float64, 5) // make([]array type, length, capacity of slice?)
    // Another way to create slices is to use the [low:high] expr:
    x_arr := []float64{ 1, 2, 3, 4, 5 }
    y_arr := x_arr[0:5] // return [1,2,3,4,5]
    // z_arr := x_arr[1:4]
    // if arr[1:4] was used it will instead return [2,3,4]
    fmt.Println(y_arr)
    fmt.Println(x_arr[1:4])

    slice_functions_append()
    slice_functions_copy()

    /* 6.3 Map Section
        A map is an unordered collection of key-value pairs. Also known as an
        associative array, a hash table or a dictionary, maps are used to look
        up a value by it's associated key. Here's an example of a map in Go:
    */
    x_map := make(map[int]int)  // initialization, make(map[key]value)
    x_map[1] = 10
    fmt.Println(x_map[1])

    // We can also delete items from a map using the built-in "delete" function
    //delete(x, 1)
    maps_example()
    maps_example_state()
}

func example_array1() {
    // this program simply gets the summation of the values inside of the array
    // then divides it by 5 = profit?
    // in short gets the average

    const numberOfGrades = 5

    var x [numberOfGrades]float64 // array x which could hold up to 5 elements; for the test scores
    x[0] = 98 // values 0 .. 4
    x[1] = 93
    x[2] = 77
    x[3] = 82
    x[4] = 83

    var total float64 = 0 // give total a type float for the result
    // could also use len(array)
    //for i := 0 ; i < numberOfGrades; i++ { // control statement for iterating through the array
    for i := 0; i < len(x) ; i++ {
        total += x[i] // every time we go through the loop the values get added up
    }
    fmt.Println(total / float64(len(x))) 
    // get the result
    // also using this method we will get an error
    // (mismatched types float64 and int)
    // solution: typecast or convert len(x) to float64
}

func example_array2() {

    const numberOfGrades = 5

    // shorter way of writing an array in Golang
    x := [5]float64{ 98, 93, 77, 82, 83 }
    /* yet another way ... damn
    x := [5]float64 {
        98,
        93,
        77,
        82,
        83,
    }
    */

    var total float64 = 0
    for _, value := range x { // using blank identifier since we don't use "i"
        total += value
    }
    fmt.Println(total / float64(len(x)));
}

func slice_functions_append() {
    // append and copy can be used to assist with slice operations
    slice1 := []int{1,2,3}
    slice2 := append(slice1, 4, 5)
    fmt.Println(slice1, slice2) // simply appends [1,2,3] with [4,5] = [1,2,3,4,5]
}

func slice_functions_copy() {
    slice1 := []int{1,2,3}
    slice2 := make([]int, 2)
    copy(slice2, slice1) // copy(to, from)
    fmt.Println(slice1, slice2)
}

func maps_example() {
    elements := make(map[string]string)
    elements["H"] = "Hydrogen"
    elements["He"] = "Helium"
    elements["Li"] = "Lithium"
    elements["Be"] = "Beryllium"
    elements["B"] = "Boron"
    elements["C"] = "Carbon"
    elements["N"] = "Nitrogen"
    elements["O"] = "Oxygen"
    elements["F"] = "Flourine"
    elements["Ne"] = "Neon"

    fmt.Println(elements["Li"]) // will return Lithuim
    //name, ok := elements["Un"] // check for the zero value in a condition 
    //fmt.Println(name, ok) // will return false 
    if name, ok := elements["Un"]; ok {
        fmt.Println(name, ok) // will print if it's a non zero value
    }
}

func maps_example_state() {
    elements := map[string]map[string]string {
        "H": map[string]string {
            "name":"Hydrogen",
            "state":"gas",
        },
        "He": map[string]string {
            "name":"Helium",
            "state":"gas",
        },
        "Li": map[string]string {
            "name":"Lithium",
            "state":"solid",
        },
        "Be": map[string]string {
            "name":"Beryllium",
            "state":"solid",
        },
        "B": map[string]string {
            "name":"Boron",
            "state":"solid",
        },
        "C": map[string]string {
            "name":"Carbon",
            "state":"solid",
        },
        "N": map[string]string {
            "name":"Nitrogen",
            "state":"gas",
        },
        "O": map[string]string {
            "name":"Oxygen",
            "state":"gas",
        },
        "F": map[string]string {
            "name":"Flourine",
            "state":"gas",
        },
        "Ne": map[string]string {
            "name":"Neon",
            "state":"gas",
        },
    }

    if el, ok := elements["Un"]; ok {
        fmt.Println(el["name"], el["state"])
    } else {
        fmt.Println("The element doesn't exist")
    }
}
