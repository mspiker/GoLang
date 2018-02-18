package main

import (
	"fmt"
	"math/cmplx"
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

// Variable declarations can be factored into blocks, just like the import statements, variables
// declared without an explicit initial value are given their zero value.  
var (
	IsAuth	bool		= false
	MaxInt	uint64		= 1<<64 - 1					// Shifts the 1 bit left 100 places, in other
													// words, the binary number that is 1 followed by
													// 63 zeros.
	CplxNo	complex128	= cmplx.Sqrt(-5 + 12i)
	aRune	rune		= 'b'
	Zero	int32
)

// Constants are just like variables, but with the const keyword.  You cannot declare constants
// using the short assignment, or := syntax.
const Pi = 3.14159267

func main() {
	
	// 1. Each Go program starts running in package main.
	// 2. Capital letters are exported names, Lowercase are not exported; public vs. private
	
	fmt.Println("01-Basics Application")
	fmt.Println(add(42, 13))
	fmt.Println(multiply(42, 13))
	
	// Example of a function returning more than one result, notice the use of := (short assignment)
	// this eliminates the need to procede the line with var like I did declaring a and b.  Short
	// assignments can only be used in a func, outsie a func every statement begins with a keyword.  
	var a, b = swap("Item 1", "Item 2")
	c, d := swap("Item 3", "Item 4")
	fmt.Println(a, b)
	fmt.Println(c, d)
	
	// Var declarations can include initializers
	var i, j int = 1, 2
	// If an initializer is present you can omit the type
	var x, y = 2, 4
	fmt.Println(i, j)
	fmt.Println(x, y)
	
	// Formatting printed values
	fmt.Printf("Type: %T Value: %v\n", IsAuth, IsAuth)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", CplxNo, CplxNo)
	fmt.Printf("Type: %T Value: %v\n", aRune, aRune)
	fmt.Printf("Type: %T Value: %v\n", Zero, Zero)
	
	fmt.Println(GetDB())
}
