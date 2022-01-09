package main

import "fmt"

// Go позволяет присваивать имена значениям, возвращаемым из функций.
// Кроме того, если в такой функции встречается оператор return без
// аргументов, то функция автоматически возвращает текущее состояние
// каждого именованного возвращаемого значения в той последовательности,
// в которой эти значения были объявлены в определении функции.

func namedMinMax(x, y int) (min, max int) {
	if x > y {
		min = y
		max = x
	} else {
		// порядок ввжен ток в определении функции -> (min, max int)
		max = y
		min = x
	}
	return
	// return min, max   -> equal to this
}

func main() {
	var y, x int
	fmt.Print("Enter value: ")
	_, err := fmt.Scanf("%d", &y)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print("Enter value: ")
	_, err = fmt.Scanf("%d", &x)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(namedMinMax(x, y))
	min, max := namedMinMax(x, y)
	fmt.Println(min, max)
	// --- использование именованных переменных очень полезно,
	// если нужно вернуть значение из функции после обработки паники
}
