package main

import (
	"fmt"
)

/* --- 1 - interface{} */

// The interface type that specifies zero methods
// is known as the empty interface:
/* interface{} */
// An empty interface may hold values of any type.
// (Every type implements at least zero methods.)

// Empty interfaces are used by code that handles values of unknown type.
// For example, fmt.Print takes any number of arguments
// of type interface{}.

func describe3(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i interface{}
	describe3(i)

	i = 42
	describe3(i)

	i = "hello"
	describe3(i)
	fmt.Println()

	/* --- 2 - TYPE ASSERTIONS */
	var inter2 interface{} = "hello"

	// A type assertion provides access to an interface value's
	// underlying concrete value  |  t := i.(T)
	s := inter2.(string)
	fmt.Println(s)
	// If i does not hold a T, the statement will trigger a panic.
	/*
		f = inter2.(float64) // panic
	*/

	// --- To test whether an interface value holds a specific type,
	// a type assertion can return two values: the underlying value
	// and a boolean value that reports whether the assertion succeeded:
	s, ok := inter2.(string)
	// If i holds a T, then t will be
	// the underlying value and ok will be true
	fmt.Println(s, ok)

	// --- If not, ok will be false
	// and t will be the ZERO VALUE of type T
	f, ok := inter2.(float64)
	fmt.Println(f, ok) // f = 0 - zero value of float64
	fmt.Println()

	/* --- 3 - TYPE SWITCHES */
	// A type switch is a construct that permits type assertions
	// in series.
	// A type switch is like a regular switch statement,
	// but the cases in a type switch specify types (not values)
	do(21)
	do("hello")
	do(true)

}

// The declaration in a type switch has the same syntax
// as a type assertion i.(T), but the specific type T
// is replaced with the keyword 'type'
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
