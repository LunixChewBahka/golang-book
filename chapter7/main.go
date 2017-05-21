package main

import (
	"fmt"
	"reflect"
)

/*
7 Functions

A function is an independent section of code that maps zero ro more input
parameters to zero or more output parameters. Functions (also known as proceduressubroutines) are often represented as a black box: (the black box represenets the function)
*/

// Functions start with a func keyword, followed by the function's name.
// in this case it's average.
// function name: average
// function parameter = 1; name: 'xs'; type:[]float64
// function signature / return type = float64
func average(xs []float64) float64 {
	//panic("Not Implemented")
	// panic is a built-in function; panic causes a runtime error
	// format: panic("Error message that you want to appear")
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func f(x int) {
	fmt.Println(x)
}

func f1() int {
	return f2()
}

func f2() (r int) {
	r = 1
	rt := reflect.TypeOf(r).Kind()
	fmt.Println("I was called by f1(). Hi, I'm from f2()")
	fmt.Printf("%T: %s\n", rt, rt)
	if rt == reflect.Int {
		fmt.Println(">> r is int")
	}
	return
}

func mrf() (int, int) {
	return 5, 6
}

// "..." you can indicate that it takes zero or more of those parameters
func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

// Even Generator, calling another function using a function; func ception?
func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func factorial(x uint) uint {
	/*
		+ Is x == 0? No. (x is 2)
		+ Find the factorial of x - 1
			+ Is x ==? No. (x is 1)
			+ Find the factorial of x - 1
				+ Is x == 0? Yes, return 1.
			+ return 1 * 1
		+ return 2 * 1
	*/

	if x == 0 {
		return 1
	}

	return x * factorial(x-1)
}

func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}

func main() {
	// invoke / make use of the function
	/*
		7.1 Your Second Function
	*/
	fmt.Println("7.1 Your Second Function")
	youCanUseAnyNameHere := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(youCanUseAnyNameHere))

	x := 5
	f(x)

	// understanding stack
	// stack = static; allocation is dealt @ compile time; fast access since it's directly stored in memory; Last In First Out approach; Thread specific;
	// heap = dynamic objects; memory allocated at runtime; accessing memory is a bit slower; heap size is limited by the size of the virtual memory; can be accessed randomly at any time; Application specific;
	f1()

	/*
		7.2 Returning Multiple Values
		Go is also capable of returning multiple values from a function
	*/
	fmt.Println("7.2 Returning Multiple Values")
	//f_x, f_y := mrf()

	/*
		7.3 Variadic Functions
		There is a special orm available for the last parameter in a Go function
	*/
	fmt.Println("7.3 Variadic Functions")
	fmt.Println(add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	/*
		This is how Println is implemented:
		func Println(a ...interface{}) (n int, err error)
	*/

	// we can also use it to pass a slice
	slc_xs := []int{1, 2, 3}
	// simple adds all the values inside the slice
	fmt.Println(add(slc_xs...))

	/*
		7.4 Closure
		It is possible to create functions inside of functions
	*/
	fmt.Println("7.4 Closure")
	closure_add := func(x, y int) int {
		return x + y
	}

	fmt.Println(closure_add(1, 1))
	cat := reflect.TypeOf(closure_add).Kind()
	fmt.Printf("%T: %s\n", cat, cat)

	inc_x := 0
	increment := func() int {
		inc_x++
		return inc_x
	}
	fmt.Println(increment())
	fmt.Println(increment())

	// func even generator call
	nextEven := makeEvenGenerator()
	even_val := 5
	for i := 0; i <= even_val; i++ {
		fmt.Println(nextEven())
	}

	/*
		7.5 Recursion
		Finally a function is able to call itself. Here is one way to compute the factorial of a number:
	*/
	fmt.Println("7.5 Recursion")
	//fmt.Println(factorial(1))
	//fmt.Println(factorial(2))
	//fmt.Println(factorial(3))
	//fmt.Println(factorial(4))
	//fmt.Println(factorial(5))
	//fmt.Println(factorial(6))
	//fmt.Println(factorial(7))
	//fmt.Println(factorial(8))
	//fmt.Println(factorial(9))
	//fmt.Println(factorial(10))
	//fmt.Println(factorial(11))
	for fact_init := uint(0); fact_init <= 5; fact_init++ {
		fact_init += 1
		// should print 11x
		fmt.Println("Call factorial inside a for-loop:", factorial(fact_init))
	}
	/*
		7.6 Defer, panic & Recover
		Go has a special statement called defer which schedules a function call to be run after the function completes. consider th e following example:
	*/
	fmt.Println("\n-----------------------")
	fmt.Println("7.6 Defer, panic & Recover")
	// defer second() // defer moves the call to second to the end of the function
	first() // should print "1st"

	/*
		usually defer is often usef when resources need to be freed in some way. For example when we open a file we need to make sure to close it later. With  defer:

		Example:
			f, _ := os.Open(filename)
			defer f.Close()
	*/

	// --- Panic and Recover
	//defer func() {
	//str := recover()
	//fmt.Println(str)
	//}()
	//panic("PANIC")
	/*
		Function Exercises
	*/
	fmt.Println("-----------------------")
	fmt.Println("Function Exercises (Answers):")
	slc_sum := []float64{100, 200, 300, 400, 500, 600, 700, 800}
	fmt.Println("Answer to number 1 is: ", answer_1(slc_sum))

	xx, yy := half(2)
	fmt.Printf("(%d, %s)\n", xx, yy)
}
