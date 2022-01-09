package main

import "fmt"

func main() {
	// Следующий код Go будет работать:
	aMap := map[string]int{}
	aMap["test"] = 1

	// Однако следующий код Go работать не будет, поскольку мы присвоили
	// значение nil хеш-таблице, которую затем пытаемся использовать:
	aMap = nil
	fmt.Println(aMap)
	aMap["test"] = 1 // panic: assignment to entry in nil map

	// Однако поиск, удаление, определение длины и использование циклов
	// range для хеш-таблицы со значением nil не приведут к краху кода.
}
