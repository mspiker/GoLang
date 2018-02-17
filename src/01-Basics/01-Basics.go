package main

import (
	"fmt"
)

// Adds two integers together and returns an integer
func add(x int, y int) int {
	return x + y
}

// Multiplys two integers and returns an integer
// When two or more consecutive named parameters share a type, omit the type from
//   all but the last parameter.  In this example x is of type int.
func multiply(x, y int) int {
	return x * y
}

// Swaps the order of the values sent into the func, a function can return any number
// of results, the type of return value is defined in the example as (string, string).
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	
	// 1. Each Go program starts running in package main.
	// 2. Capital letters are exported names, Lowercase are not exported; public vs. private
	
	fmt.Println("01-Basics Application")
	fmt.Println(add(42, 13))
	fmt.Println(multiply(42, 13))
	fmt.Println(swap("First Item", "Last Item"))
}
