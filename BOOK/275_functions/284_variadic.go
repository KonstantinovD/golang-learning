package main

import (
	"fmt"
	"os"
)

func varFunc(input ...string) {
	fmt.Println(input)
}

// В функции с переменным числом аргументов оператор упаковки может
// использоваться только один раз
func oneByOne(message string, s ...int) int {
	fmt.Println(message)
	sum := 0
	for index, a := range s {
		fmt.Println(index, ") ", a)
		sum = sum + a
	}
	s[0] = -1000
	return sum
}

func main() {
	arguments := os.Args
	varFunc(arguments[1:]...)
	sum := oneByOne("Adding numbers...", 1, 2, 3, -4, 5, 5)
	fmt.Println("Sum:", sum)
	fmt.Println()
	s := []int{1, 2, 3}
	sum = oneByOne("Adding numbers...", s...)
	fmt.Println("modified arr:", s)
}
