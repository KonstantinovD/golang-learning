package main // определение пакета для текущего файла

import ( // подключение пакета fmt
	"fmt"
	"strconv"
)

func main() {
	//объява переменнной: var имя_переменной тип_данных
	var hello string
	hello = "Hello world"
	//нельзя не использовать эту переменную: hello declared but not used
	fmt.Println(hello)

	// списочное объявление переменных
	var a, b, c int
	a = 11
	b = 12
	c = 13
	var k = a + b + c
	fmt.Println(k)

	// Если мы хотим сразу определить несколько переменных и присвоить
	// им начальные значения, то можно обернуть их в скобки:
	var (
		name string = "Tom"
		age  int    = 27
	)
	fmt.Println(name + " - " + strconv.Itoa(age))

	name = "Billy"
	fmt.Println("changed name: " + name)

	// Также допустимо краткое определение переменной в формате:
	secondName := "Bons"
	//Тип данных выводится автоматически из присваиваемого значения.
	fmt.Println("second name: " + secondName)

	/*** КОНСТАНТЫ ***/
	const sq2 float64 = 1.414
	const (
		pi float64 = 3.1415
		e  float64 = 2.7182
	)
	const n = 5 // тип int - выводится неявно по значению

	// Если определяется последовательность констант, то инициализацию
	// значением можно опустить для всех констант, кроме первой. В этом
	// случае константа без значения получит значение предыдущей константы:
	const (
		c1 = 1
		c2
		c3
		c4 = 3
		c5
	)
	fmt.Println(c1, c2, c3, c4, c5) // 1, 1, 1, 3, 3

	fmt.Print("Press 'Enter' to continue...")
	fmt.Scanln() // wait for Enter Key

	// Типизация констант накладывает доп. ограничения на них
	const s1 = 123
	const s2 float64 = 123
	var v1 float32 = s1 * 12
	//var v2 float32 = s2 * 12  - // isn't compiling
	fmt.Println(v1)
	//fmt.Println(v2)
}
