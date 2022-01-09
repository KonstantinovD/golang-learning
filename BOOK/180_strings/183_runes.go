package main

import "fmt"

// Руны — это значения типа int32, следовательно, тип Go, который
// используется для представления символов Unicode

func main() {
	// Рунный литерал — это символ, заключенный в одинарные кавычки.
	// Рунный литерал также можно рассматривать как рунную константу
	const r1 = '€'
	fmt.Println("(int32) r1:", r1)   // (int32) r1: 8364
	fmt.Printf("(HEX) r1: %x\n", r1) // (HEX) r1: 20ac
	fmt.Printf("(as a String) r1: %s\n", r1)
	// r1: %!s(int32=8364)
	// Чтобы преобразовать руну в символ, нужно использовать %c.
	fmt.Printf("(as a character) r1: %c\n", r1)
	// (as a character) r1: €

	// байтовый срез представляет собой набор рун и вывод байтового среза
	// с помощью fmt.Println() дает срез числовых значений рун
	fmt.Println("A string is a collection of runes:",
		[]byte("Mihalis"))
	// A string is a collection of runes: [77 105 104 97 108 105 115]
	aString := []byte("Mihalis")
	for x, y := range aString {
		fmt.Println(x, y)
		fmt.Printf("Char: %c\n", aString[x])
	}
	// Для вывода байтового среза как строку используется %s.
	fmt.Printf("%s\n", aString)
}
