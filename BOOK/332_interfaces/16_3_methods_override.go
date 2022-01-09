package main

import (
	"fmt"
)

type MyThinger struct{}

func (me *MyThinger) Foo() int {
	return 1
}

func (me *MyThinger) Bar() int {
	return 2
}

type userThinger struct {
	// Composition!
	MyThinger
}

func NewUserThinger() *userThinger {
	return &userThinger{
		MyThinger: MyThinger{},
	}
}

// "overriding"
func (me *userThinger) Foo() int {
	return 42
}

func main() {
	// --- также доступен ОВЕРРАЙДИНГ МЕТОДОВ
	// вариант 1 - без интерфейсов:
	thinger := NewUserThinger()
	fmt.Println(thinger.Bar())
	fmt.Println(thinger.Foo())
	fmt.Println()

	skyCrapper := NewSkyCraper()
	fmt.Println(skyCrapper.Bar())
	fmt.Println(skyCrapper.Foo())
	fmt.Println()
}

type Building interface {
	Foo() int
	Bar() int
}

func NewBuilding() Building {
	return &House{}
}

type House struct{}

func (me *House) Foo() int {
	return 1
}

func (me *House) Bar() int {
	return 2
}

type SkyCraper struct {
	// Composition!
	Building
}

func NewSkyCraper() Building {
	return &SkyCraper{
		Building: NewBuilding(),
	}
}

// "overriding"
func (me *SkyCraper) Foo() int {
	return 42
}
