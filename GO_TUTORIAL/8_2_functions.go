package main // определение пакета для текущего файла

import ( // подключение пакета fmt
	"fmt"
)

/* ТИП ФУНКЦИИ */
// Каждая функция имеет определенный тип, который складывается из
//списка типов параметров и списка типов возвращаемых рехультатов.
func add(x int, y int) int {
	return x + y
}

// Эта функция представляет тип 'func(int, int) int'.
// Этому же типу будет соответствовать следующая функция:
func multiply(x int, y int) int {
	return x * y
}

// Имена различаются, но по типу параметров и по типу возвращаемого результата
// она соответствует вышеуказанному типу функции.
func display(message string) {
	fmt.Println(message)
}

// Эта функция имеет тип func(string)

// Это значит, что мы можем определять переменные или параметры функций,
// которые будут представлять определенный тип функциии.
// То есть фактически переменная может быть функцией. Например:

func main() {
	var f func(int, int) int = add
	fmt.Println(f(3, 4))
	fmt.Println()

	var x = f(4, 5) // 9
	fmt.Println(x)

	f = multiply // теперь переменная f указывает на функцию multiply
	fmt.Println(f(3, 4))
	fmt.Println()

	// функция может передаваться в качестве параметра в другую функцию:
	if action(10, 20, add) == action(5, 6, multiply) {
		fmt.Println("30 equal 30")
	}
	fmt.Println()

	// Функция как результат другой функции
	f2 := selectFn(1)
	fmt.Println(f2(3, 4)) // 7

	f2 = selectFn(3)
	fmt.Println(f2(3, 4)) // 12
	fmt.Println()

}

func action(n1 int, n2 int, operation func(int, int) int) int {
	result := operation(n1, n2)
	return result
}

func selectFn(n int) func(int, int) int {
	if n == 1 {
		return add
	} else if n == 2 {
		return nil
	} else {
		return multiply
	}
}
