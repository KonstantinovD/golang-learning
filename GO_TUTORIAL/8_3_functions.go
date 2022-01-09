package main // определение пакета для текущего файла

import ( // подключение пакета fmt
	"fmt"
	"strconv"
)

func main() {
	// Анонимные функции - это функции, которым не назначено имя.
	// Они могут определяться внутри других функций
	// и иметь доступ к контексту выполнения.
	f := func(x, y int) int { return x + y }
	fmt.Println(f(3, 4)) // 7
	fmt.Println(f(6, 7)) // 13
	fmt.Println()

	//Анонимная функция как аргумент функции
	action(10, 25, func(x int, y int) int { return x + y }) // 35
	action(5, 6, func(x int, y int) int { return x * y })   // 30
	fmt.Println()

	// Анонимная функция может быть результатом друой функции
	f2 := messageProcessor("DEFAULT")
	fmt.Println(f2("NORMAL", 14))
	f2 = messageProcessor("ERROR")
	fmt.Println(f2("ERROR", 15))
	fmt.Println()

	// Преимуществом анонимных функций является то, что они
	// имеют доступ к окружению, в котором они вызываются.
	// Такие функции называются замыканиями (closures)
	f3 := square()
	f4 := square()
	fmt.Println("f3: ", f3()) // 9
	fmt.Println("f3: ", f3()) // 16
	fmt.Println("f3: ", f3()) // 25
	fmt.Println("f4: ", f4()) // 9
	fmt.Println("f3: ", f3()) // 36
	fmt.Println("f4: ", f4()) // 16

	// Рекурсия, куда же без нее)))
	fmt.Println()
	fmt.Println("recusrion")
	fmt.Println()
	fmt.Println("factorials")
	fmt.Println(factorial(4)) // 24
	fmt.Println(factorial(5)) // 120
	fmt.Println(factorial(6)) // 720

	fmt.Println("fibbonachi")
	fmt.Println(fibbonachi(4)) // 3
	fmt.Println(fibbonachi(5)) // 5
	fmt.Println(fibbonachi(6)) // 8
}

func action(n1 int, n2 int, operation func(int, int) int) {
	result := operation(n1, n2)
	fmt.Println(result)
}

func messageProcessor(msg string) func(string, int) string {
	switch msg {
	case "DEFAULT":
		{
			return func(s string, n int) string {
				return "System in a " + s + " state. Value = " + strconv.Itoa(n)
			}
		}
	case "WARNING":
		{
			return func(s string, n int) string {
				return "Warning 01_message. Value = " + strconv.Itoa(n)
			}
		}
	default:
		return func(s string, n int) string {
			return "System cannot process 01_message"
		}
	}

}

func square() func() int {
	var x int = 2
	return func() int {
		x++
		return x * x
	}
}

func factorial(n uint) uint {

	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func fibbonachi(n uint) uint {

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibbonachi(n-1) + fibbonachi(n-2)
}
