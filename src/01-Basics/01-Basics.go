package main

import (
	"fmt"
	"math/cmplx"
	"runtime"
	"time"
)

// General Information about Go
// ----------------------------
// 1. Each Go program starts running in package main.
// 2. Capital letters are exported names, Lowercase are not exported; public vs. private


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

func ForLoops() {
	sum := 0
	
	// The basic for loop has three components separated by semicolons.
	//   - the init statement: executed before the first iteration
	//        - variables declared are visible only in the scope of the for statement
	//   - the condition expression: evaluated before every iteration
	//   - the post statement: executed at the end of every iteration
	// Unlike other languages, no parenthesis surround the three components, braces { } are always required.
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Printf("The sum after the for loop is %d\n", sum)
	
	// The init and post statements are optional
	sumA := 1
	for ; sumA < 1000; {
		sumA += sumA
	}
	fmt.Printf("The sum after the second for loop is %d\n", sumA)
	
	// For is also While, just drop the semicolons 
	sumB := 1
	for sumB < 1000 {
		sumB += sumB
	}
	fmt.Printf("The sum after the while loop is %d\n", sumB)

	// See the func ForeverLoops for an example of an infinite loop.
	
}

func ForeverLoops() {
	// Omit the loop condition and an infinite loop is expressed.
	for {

	}
}

func IfLogic() {
	// Like for loops, logic in an if does not need parenthesis like other languages.
	x := 2
	if x < 0 {
		fmt.Println("x was less than 2")
	}
	
	// If statements can contain short statements which execute before the condition.  Here 
	// we set v to 2 times that of x then check to see if it is greater than 3 which it should be.
	// Variables, like v, declared by the statement are only in scope until the end of the if/else.  
	// In this example we also demonstrate the use of else.  
	if v := x * 2 ; v > 3 {
		fmt.Printf("v is %d so we made it in the if statement.", v)
	} else {
		fmt.Printf("v is greater than 3 becuase it is %d", v)
	}
}

func UsingSwitch() {
	// Switch is a shorter way to write a sequence of if statements.  The difference between 
	// Go and other languages is that it only runs the selected case and not all cases that follow
	// as other languages do.  A break statement is NOT required in Go, it takes care of that.  
	// Another difference is that the case does not need to be a constant.  The order of evaluation
	// if from top to bottom, stopping when a case succeeds so if the OS were "darwin" it would stop
	// there and go no further.  
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
		case "darwin":
			fmt.Println("OS X.")
		case "linux":
			fmt.Println("Linux.")
		default: 
			// all other possible OS
			fmt.Printf("%s.", os)
	}
	
	// Switch without a condition is the same as "switch true".  This construct can be a clean
	// way to write long if-then-else chains.
	t := time.Now()
	switch {
		case t.Hour() < 12: 
			fmt.Println("Good morning!")
		case t.Hour() < 17:
			fmt.Println("Good afternoon.")
		default: 
			fmt.Println("Good evening.")
	}
}

func main() {
	
	
	
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
	
	// Formatting printed values, notice I am using Printf, not Println
	fmt.Printf("Type: %T Value: %v\n", IsAuth, IsAuth)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", CplxNo, CplxNo)
	fmt.Printf("Type: %T Value: %v\n", aRune, aRune)
	fmt.Printf("Type: %T Value: %v\n", Zero, Zero)
	
	// This calls the func GetDB which lives in another file, see 02-MongoDb.go.  You will notice
	// that the package name is the same and 02-MongoDb has no func call main().  This is much like
	// namespaces in other languages.  
	fmt.Println(GetDB())
	
	// This func demonstrates For and While
	ForLoops()
	
	// This func demonstrates If and Else
	IfLogic()
	
	// This func demonstrates Switch
	UsingSwitch()
	
	// Deferring execution until the surrounding function returns.  
	defer fmt.Print("world. This was an example of the defer statemnt.")
	fmt.Print("hello ")
	
	
	
}
