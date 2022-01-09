package main

import "fmt"

/* ___PANIC___ & ___REECOVER___ */
// Panic func stops the ordinary flow of control and begins panicking.
// When the function F calls panic, execution of F stops, any deferred
// functions in F are executed normally, and then F returns to its
// caller. To the caller, F then behaves like a call to panic. The
// process continues up the stack until all functions in the current
// goroutine have returned, at which point the program crashes. Panics
// can be initiated by invoking panic directly. They can also be caused
//by runtime errors, such as array out-of-bounds.

// Recover is a built-in function that regains control of a panicking
// goroutine. Recover is only useful inside deferred functions. During
// normal execution, a call to recover will return nil and have no other
// effect. If the current goroutine is panicking, a call to recover will
// capture the value given to panic and resume normal execution.

func f() {
	// panic is called with an argument, and it is the value
	// which 'recover()' func returns
	defer func() {
		if r := recover(); r != nil {
			// r == i where panic(i)
			fmt.Println("Recovered in f", r)
		}
	}()
	// or 	de_f := func() {...}
	//      defer de_f()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(i)
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}

func h() {
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func main() {
	f()
	fmt.Println("Returned normally from f.")

	fmt.Println()
	fmt.Println("........................")
	fmt.Println("panic without recover - program will crashes")
	fmt.Println()
	h()
}
