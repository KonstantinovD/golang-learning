package main

import "fmt"

func main() {
	a6 := []int{1, 2, 3, 4, 5, 6}
	a4 := []int{11, 22, 33, 44}
	fmt.Println("a6:", a6)
	fmt.Println("a4:", a4)
	copy(a6, a4)
	fmt.Println("a6:", a6)
	fmt.Println("a4:", a4)
	fmt.Println("---------------------")
	fmt.Println()
	// --- Поскольку в a6 больше элементов, чем в a4, все элементы a4
	// будут скопированы в a6. Однако, поскольку у a4 только четыре
	// элемента, а у a6 — шесть элементов, последние два элемента a6
	// останутся прежними.

	b6 := []int{11, 22, 33, 44, 55, 66}
	b4 := []int{1, 2, 3, 4}
	fmt.Println("b6:", b6)
	fmt.Println("b4:", b4)
	copy(b4, b6)
	fmt.Println("b6:", b6)
	fmt.Println("b4:", b4)
	fmt.Println("---------------------")
	fmt.Println()
	// В этом случае в b4 будут скопированы только первые четыре
	// элемента b6, так как емкость b4 составляет всего четыре элемента.

	array4 := [4]int{1, 2, 3, 4}
	s6 := []int{11, 22, 33, 44, 55, 66}
	fmt.Println("array4:", array4)
	fmt.Println("s6:", s6)
	copy(s6, array4[0:])
	fmt.Println("array4:", array4[:]) // better - передача ток слайса
	// а не всего массива целиком
	fmt.Println("s6:", s6)
	fmt.Println("---------------------")
	fmt.Println()
	// Массив из четырех элементов копируется в срез из шести элементов.
	// массив преобразуется в срез с помощью нотации [:]

	array5 := [5]int{1, 2, 3, 4, 5}
	s7 := []int{11, 22, 33, 44, 55, 66, 77}
	fmt.Println("array5:", array5)
	fmt.Println("s7:", s7)
	copy(array5[0:], s7)
	fmt.Println("array5:", array5)
	fmt.Println("s7:", s7)
	// Здесь показано, как скопировать срез в массив,
	// в котором есть место для пяти элементов
}
