package main

import "fmt"

// NOTE: Functions in GO.
// * Functions are first class citizen in GO.
// * Every Go programs start with the main function
// * without parameters.
// * The unique part is we can return multiple values from the Go functions.

// Example: GO functions
func div(num int, denom int) int {
	if denom == 0 {
		return 0
	}
	return num / denom
}

// NOTE: When we have two or more parameters in function with same type, we can specify type once for all of them like this:
func moreParams(param1, param2, param3, param4, param5 int) {}

func main() {
	result := div(5, 2)
	fmt.Println("result of div: ", result)
}
