package main

import (
	"fmt"
)

// --- 1. Мы можем вызвать метод интерфейса (и метод класса),
// если сам класс <nil>

type I interface {
	f() // the own function of interface
}

func describe2(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

type T struct {
	S string
}

func (t *T) f() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func (t *T) own() {
	if t == nil {
		fmt.Println("error: <nil>")
		return
	}
	fmt.Println("value: " + t.S)
}

func main() {
	var i I

	var t *T
	i = t
	i.f()
	// --- CAST INTERFACE to concrete types
	tobject, ok := i.(*T)
	if ok {
		tobject.own() // should be "error: <nil>" string printed
		// because tobject points to <nil> yet
	}
	describe2(i)
	fmt.Println()

	i = &T{"hello"}
	i.f()
	tobject, ok = i.(*T)
	if ok {
		tobject.own()
	}
	describe2(i)
	fmt.Println()

	var i2 I
	describe2(i2)
	// --- 2. A nil interface value holds neither value nor concrete type
	// Calling a method on a nil interface is a runtime error: there is
	// no type inside the interface tuple to indicate
	// which concrete method to call.

	/* i2.f() */
	//! ERROR
	//panic: runtime error: invalid memory address or nil pointer dereference
}
