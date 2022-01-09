package main

import "fmt"

func main() {
	aSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(aSlice)
	integers := make([]int, 2)
	fmt.Println(integers)
	integers = nil // set slice to []
	fmt.Println(integers)

	anArray := [5]int{-1, -2, -3, -4, -5}
	refAnArray := anArray[:]

	fmt.Println(anArray)
	fmt.Println(refAnArray)
	anArray[4] = 100
	fmt.Println(refAnArray)

	s := make([]byte, 5)
	fmt.Println(s)
	twoD := make([][]int, 3)
	fmt.Println(twoD)
	fmt.Println()
	// Поскольку срезы в Go инициализируются автоматически, все элементы
	// этих двух срезов будут иметь нулевое значение, соответствующее
	// типу данного среза, которое для целых чисел равно 0,
	// а для срезов — nil, т е []

	for i := 0; i < len(twoD); i++ {
		for j := 0; j < 2; j++ {
			twoD[i] = append(twoD[i], i*j)
		}
	}

	for _, x := range twoD {
		for i, y := range x {
			fmt.Println("i:", i, "value:", y)
		}
		fmt.Println()
	}
}
