package main

import (
	"fmt"
	"os"
	"reflect"
)

type t1 int
type t2 int

type c struct {
	X    int
	Y    float64
	Text string
}

// Цель функции compareStruct — выяснить, являются ли две переменные
// типа a абсолютно одинаковыми.
func (c1 c) compareStruct(c2 c) bool {
	r1 := reflect.ValueOf(&c1).Elem()
	r2 := reflect.ValueOf(&c2).Elem()
	for i := 0; i < r1.NumField(); i++ {
		if r1.Field(i).Interface() != r2.Field(i).Interface() {
			return false
		}
	}
	return true
}

// выводит список методов переменной
func printMethods(i interface{}) {
	r := reflect.ValueOf(i)
	t := r.Type()
	fmt.Printf("Type to examine: %s\n", t)

	for j := 0; j < r.NumMethod(); j++ {
		m := r.Method(j).Type()
		fmt.Println(t.Method(j).Name, "-->", m)
	}
}

func main() {
	x1 := t1(100)
	x2 := t2(100)
	fmt.Printf("The type of x1 is %s\n", reflect.TypeOf(x1))
	fmt.Printf("The type of x2 is %s\n", reflect.TypeOf(x2))
	fmt.Println()

	var p struct{}
	r := reflect.New(reflect.ValueOf(&p).Type()).Elem()
	fmt.Printf("The type of r is %s\n", reflect.TypeOf(r))
	fmt.Println()

	c1 := c{1, 2.1, "C1"}
	c2 := c{1, -2, "C2"}
	c3 := c{1, 2.1, "C1"}
	if c1.compareStruct(c1) {
		fmt.Println("Equal!")
	}
	if !c1.compareStruct(c2) {
		fmt.Println("Not Equal!")
	}
	if c1.compareStruct(c3) {
		fmt.Println("Equal!")
	}
	fmt.Println()

	fmt.Println("Methods of 'os.File':")
	var f *os.File
	printMethods(f)
}
