package main

import "fmt"

func funReturnFun() func() int {
	i := 0
	return func() int {
		i++
		return i * i
	}
}

func main() {
	i := funReturnFun()
	j := funReturnFun()

	fmt.Println("i1:", i()) // i1: 1
	fmt.Println("i2:", i()) // i2: 4
	fmt.Println("j1:", j()) // j1: 1
	fmt.Println("j2:", j()) // j2: 4
	fmt.Println("i3:", i()) // i3: 9
}
