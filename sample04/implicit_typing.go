package main // определение пакета для текущего файла

import ( // подключение пакета fmt
	"fmt"
	"reflect"
)

func main() {

	/*** ЗНАЧЕНИЕ ПО УМОЛЧАНИЮ ***/
	// Если переменной не присвоено значение, то она имеет значение
	// по умолчанию, которое определено для ее типа. Для числовых типов
	// это число 0, для логического типа - false, для строк - ""
	var defvalue int32
	fmt.Println("Int default value:", defvalue)
	var floatdefvalue float64
	fmt.Println("float64 default value:", floatdefvalue)
	var booldefvalue bool
	fmt.Println("Bool default value:", booldefvalue)
	var strdefvalue string
	fmt.Println("String default value:", strdefvalue)
	fmt.Println()

	/*** НЕЯВНАЯ ТИПИЗАЦИЯ ***/

	// При определении переменной можно опускать тип, если мы явно
	// инициализируем переменную каким-нибудь значением
	var name = "Tom"

	//То же самое происходит при кратком определении переменной
	secondName := "Blader"

	// Неявная типизация нескольких переменных:
	var (
		value  = "A"
		number = 27
	)
	var town, year = "Minsk", 1067

	fmt.Println(name, secondName, value, number, town, year)
	fmt.Println(
		reflect.TypeOf(name),
		reflect.TypeOf(secondName),
		reflect.TypeOf(value),
		reflect.TypeOf(number),
		reflect.TypeOf(town),
		reflect.TypeOf(year))

	fmt.Print("Press 'Enter' to continue...")
	fmt.Scanln()
}
