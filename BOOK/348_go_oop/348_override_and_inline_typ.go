package main

import (
	"fmt"
)

type first struct{}

func (a first) F() {
	a.shared()
}

func (a first) shared() {
	fmt.Println("This is shared() from first!")
}

type second struct {
	first
}

func (a second) shared() {
	fmt.Println("This is shared() from second!")
}

func main() {
	first{}.F()       // This is shared() from first!
	second{}.shared() // This is shared() from second!
	fmt.Println()

	i := second{}
	i.shared()       // This is shared() from second!
	i.first.F()      // This is shared() from first!
	i.first.shared() // This is shared() from first!
	i.F()            // This is shared() from first! = i.first.F()
	fmt.Println()
	// Так происходит потому, что это отношения 'has-a', а не 'is-a'

	// включение анонимного поля класса - не создает отношений 'is-a'
	var g interface{} = i
	switch g.(type) { // Not 'first' - выводимый результат
	case first:
		fmt.Println("'first'")
	default:
		fmt.Println("Not 'first'")
	}

	j := first{}
	var fsArr []first = make([]first, 2)
	fsArr[0] = j
	//fsArr[1] = i  - ERROR - not type first - 'has-a' отношения
	fsArr[1] = i.first
	fmt.Println(fsArr[0], fsArr[1])
}
