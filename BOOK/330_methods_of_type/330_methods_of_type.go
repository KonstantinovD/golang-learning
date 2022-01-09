package main

import "fmt"

type twoInts struct {
	X int64
	Y int64
}

type customInt int64

func regularFunction(a, b twoInts) twoInts {
	temp := twoInts{X: a.X + b.X, Y: a.Y + b.Y}
	return temp
}

// Метод типа (type method) в Go — это функция со специальным
// аргументом-приемником. Такой метод объявляется как обычная функция
// с дополнительным параметром, который ставится перед именем функции.
// Данный параметр связывает функцию с типом этого дополнительного
// параметра. Именно этот параметр называется приемником метода.
func (a twoInts) method(b twoInts) twoInts {
	temp := twoInts{X: a.X + b.X, Y: a.Y + b.Y}
	return temp
}

func (a customInt) doubleInt() int64 {
	return int64(a * 2)
}

func main() {
	i := twoInts{X: 1, Y: 2}
	j := twoInts{X: -5, Y: -2}
	fmt.Println(regularFunction(i, j))
	fmt.Println(i.method(j))

	ci := customInt(14)
	fmt.Println("custom int:", ci,
		"| doubled custom int:", ci.doubleInt())
}
