package main // определение пакета для текущего файла

import ( // подключение пакета fmt
	"fmt"
	"strconv"
)

func main() {

	/*** 01 - ЦЕЛЫЕ ЧИСЛА ***/

	var a int8 = -1
	// int8 представляет целое число от -128 до 127
	// занимает в памяти 1 байт (8 бит)
	var b uint8 = 2 // целое число от 0 до 255
	var c byte = 3  // byte - синоним типа uint8

	var d int16 = -4 // целое число от -32768 до 32767
	var e uint16 = 5 // целое число от 0 до 65535

	var f int32 = -1 // целое число от -2147483648 до 2147483647
	var g rune = -1  // синоним типа int32
	var h uint32 = 1 // целое число от 0 до 4294967295

	var j int64 = -1
	var k uint64 = 1

	var l int = 1 // представляет целое число со знаком,
	//которое в зависимости о платформы может занимать либо 4 байта,
	// либо 8 байт. То есть соответствовать либо int32, либо int64.
	var m uint = 1

	fmt.Printf("%d %d %d %d %d %d %d %d %d %d %d %d \n",
		a, b, c, d, e, f, g, h, j, k, l, m)

	/*** 02 - ЧИСЛА С ПЛАВАЮЩЕЙ ТОЧКОЙ ***/

	var n float32 = 18
	var o float32 = 4.5 // представляет число с плавающей точкой
	// от 1.4*(10^(-45)) до 3.4*10^38(для положительных). В памяти - 4 байта
	// Тип float32 обеспечивает шесть десятичных цифр точности

	var p float64 = 0.23
	var pi float64 = 3.14
	var exponent float64 = 2.7 // представляет число с плавающей точкой
	// от4.9*10^(-324) до 1.8*10^(308) (для положительных). В памяти - 8 байт
	// Тип float64 обеспечивает около 15 десятичных цифр точности

	fmt.Printf("%f %f %f %f %f\n", n, o, p, pi, exponent)

	/*** 03 - КОМПЛЕКСНЫЕ ЧИСЛА ***/

	var complex1 complex64 = 1 + 2i //комплексное число, где
	// вещественная и мнимая части представляют числа float32
	var complex2 complex128 = 4 + 3i // а здесь - float64

	fmt.Print(strconv.FormatComplex(complex128(complex1), 'f', 0, 64) + " ")
	fmt.Println(strconv.FormatComplex(complex2, 'e', 0, 64))

	/*** 04 - ТИП BOOLEAN ***/

	var isAlive bool = true
	var isEnabled bool = false
	fmt.Println(isAlive && isEnabled, isAlive || isEnabled)
	// var isDigit bool = 0  -  only true/falsr

	/*** 05 - ТИП STRING ***/

	var name string = "Том Сойер"
	fmt.Println(name)

	/*** 06 - ДЕЛЕНИЕ ***/
	var div1 float32 = 10 / 4.5
	fmt.Println(div1)
	var div2 float32 = 11.5 / 4
	fmt.Println(div2)
	var div3 float32 = 10 / 4 // = 2
	fmt.Println(div3)
	var div4 float32 = 10 / float32(4)
	fmt.Println(div4)

	fmt.Print("Press 'Enter' to continue...")
	fmt.Scanln()
}
