package main

import (
	"fmt"
)

// see https://blog.golang.org/defer-panic-and-recover

/* ___DEFER___ */
// A defer statement pushes a function call onto a list. The list of
//saved calls is executed after the surrounding function returns.
//Defer is commonly used to simplify functions that perform
//various clean-up actions.

// Если несколько функций вызываются с оператором defer, то те функции,
// которые вызываются раньше, будут выполняться позже всех.
func runDefer() {
	defer fmt.Println("defer 1")
	fmt.Println("start execution")
	defer fmt.Println("defer 2")
	fmt.Println("finish execution")
}

////////// Существует 3 основных правила: ////////////

// _1_  Аргументы отложенной функции принимают значения которые были на
// момент объявления функции с оператором defer
func a() {
	i := 0
	defer fmt.Println(i) // напишет 0, а не 1
	i++
}

// _2_ Deferred function calls are executed in Last In First Out (LIFO)
// order after the surrounding function returns.
func b() { // This function prints "3210"
	for i := 0; i < 4; i++ {
		defer fmt.Print(i)
	}
}

// _3_ Deferred functions may read and assign
// to the returning function's named return values.
func c() (i int) {
	// deferred function increments the return value i after
	// the surrounding function returns. Thus, this function returns 2
	defer func() { i++ }()
	return 1
}

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
	runDefer()
	fmt.Println()
	a()
	fmt.Println()
	b()
	fmt.Println()
	fmt.Println()
	fmt.Println("'return 1' written, but 2 ll'be returned due to defer: ", c())
	fmt.Println()
	fmt.Println("........................")
	fmt.Println()

	f()
	fmt.Println("Returned normally from f.")

	fmt.Println()
	fmt.Println("........................")
	fmt.Println("panic without recover - program will crashes")
	fmt.Println()
	h()
}
