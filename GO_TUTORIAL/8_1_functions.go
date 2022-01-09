package main // определение пакета для текущего файла

import ( // подключение пакета fmt
	"fmt"
)

func main() {

	/* func имя_функции (список_параметров) (типы_возвращаемых_значений){
		выполняемые_операторы
	} */
	// В скобках идет список параметров. После списка параметров
	// определяются типы возвращаемых из функции значений

	hello()
	fmt.Println()

	add(4, 5) // x + y = 9
	fmt.Println()
	// Если несколько параметров подряд имеют один и тот же тип, то мы можем
	// указать тип только для последнего параметра, а предыдущие параметры
	// также будут представлять этот тип:
	addMany(1, 2, 3.4, 5.6, 1.2)
	fmt.Println()

	// аргументы в функцию всегда передаются по значению
	var a = 6
	fmt.Println("a before: ", a)
	increment(a)
	fmt.Println("a after: ", a)
	fmt.Println()
	// В Go функция может принимать неопределенное количество параметров
	// одного типа. Например, нам надо получить сумму чисел:
	fmt.Println(addList(1, 2, 3))    // sum = 6
	fmt.Println(addList(1, 2, 3, 4)) // sum = 10
	fmt.Println(addList(5, 6, 7, 2, 3))
	fmt.Println()
	// От этого случае следует отличать передачу среза в качестве параметра:
	/* addList([]int{1, 2, 3}) */
	// ОШИБКА ! так как передача среза не эквивалентна передаче
	// неопределенного количества параметров того же типа

	// При передаче среза надо указать после аргумента-массива многоточие
	fmt.Println(addList([]int{5, 6, 7}...))
	var nums = []int{5, 6, 7, 2, 3}
	fmt.Println(addList(nums...))
	fmt.Println()

	// Именованные возвращаемые результаты
	fmt.Println("11 + 12 =", addAndReturnNamed(11, 12), " (named)")
	fmt.Println()

	// Возвращение нескольких значений
	// В этом случае после списка параметров указывается в скобках список
	// типов возвращаемых значений:
	var age, name = returnTwo(4, 5, "Tom", "Simpson")
	fmt.Println(age, "|", name) // 9 Tom Simpson
	// Альтернативный способ передачи переменным результатов функции
	// И также в даном случае можно использовать именованные результаты
	_, name2 := returnTwoNamed(4, 5, "Tom", "Simpson")
	fmt.Println("fullname again: " + name2)
	fmt.Println()

}

func hello() {
	fmt.Println("Hello World")
}

func add(x int, y int) {
	var z = x + y
	fmt.Println("x + y = ", z)
}

func addMany(x, y int, a, b, c float32) {
	var z = x + y
	var d = a + b + c
	fmt.Println("x + y = ", z)
	fmt.Println("a + b + c = ", d)
}

func increment(x int) {
	fmt.Println("x before: ", x)
	x = x + 20
	fmt.Println("x after: ", x)
}

func addList(numbrs ...int) int {
	var sum = 0
	for _, number := range numbrs {
		sum += number
	}
	return sum
}

func addAndReturnNamed(x, y int) (z int) {
	z = x + y
	return // instead of 'return z' we can just wwrite 'return'
}

func returnTwo(x, y int, firstName, lastName string) (int, string) {
	var z int = x + y
	var fullName = firstName + " " + lastName
	return z, fullName
}

func returnTwoNamed(x, y int, firstName, lastName string) (z int, fullName string) {
	z = x + y
	fullName = firstName + " " + lastName
	return // instead of 'return z, fullName'
}
