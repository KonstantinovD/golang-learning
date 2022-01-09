package main

import (
	"fmt"
	"math"
)

// кортеж — это конечный упорядоченный список, состоящий
// из нескольких частей. Go не поддерживает тип кортежа и,
// следовательно, официально не обслуживает кортежи, несмотря на то
// что поддерживает способы использования кортежей
func retFour(x int) (int, int, int, int) {
	k := (int)(math.Pow(2, float64(x)))
	return 2 * x, x * x, k, -x
}

func main() {
	fmt.Print("retFour(5): ")
	fmt.Println(retFour(5))
	n1, n2, n3, n4 := retFour(10)
	fmt.Print("retFour(10): ")
	fmt.Println(n1, n2, n3, n4)
	n1, n2 = n2, n1
	n3, n4 = n4, n2
	fmt.Print("after mix: ")
	fmt.Println(n1, n2, n3, n4)
}
