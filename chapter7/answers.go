package main

// import "fmt"

// could also be func sum(myVar... int)
func answer_1(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total
}

func half(x int) (int, string) {
	// process result
	result := x / 2

	// trap conditionals
	if x == 1 {
		return 0, "false"
	} else if result%2 == 0 {
		return result, "true"
	} else {
		return result, "false"
	}
}
