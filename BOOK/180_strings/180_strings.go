package main

import "fmt"

func main() {
	// --- существуют тонкие различия между символами, рунами и байтами,
	// а также различия между строками и строковыми литералами
	// Строка Go — это байтовый срез, предназначенный только для чтения,
	// который содержит байты любого типа и имеет произвольную длину.
	//Строковый литерал можно определить следующим образом:
	const sLiteral = "\x99\x42\x32\x55\x50\x35\x23\x50\x29\x9c"
	fmt.Println(sLiteral)           // �B2UP5#P)�
	fmt.Printf("x: %x\n", sLiteral) // x: 9942325550352350299c
	fmt.Printf("Literal length: %d\n", len(sLiteral))
	// Literal length: 10

	for i := 0; i < len(sLiteral); i++ {
		fmt.Printf("%x ", sLiteral[i])
		// 99 42 32 55 50 35 23 50 29 9c
	}
	fmt.Println()

	// вывод строки в двойных кавычках
	fmt.Printf("q: %q\n", sLiteral) // q: "\x99B2UP5#P)\x9c"
	// вывод только символов ASCII
	fmt.Printf("+q: %+q\n", sLiteral) // +q: "\x99B2UP5#P)\x9c"
	// появление пробелов между выводимыми байтами
	fmt.Printf(" x: % x\n", sLiteral)
	// x: 99 42 32 55 50 35 23 50 29 9c
	fmt.Printf("s: As a string: %s\n", sLiteral)
	// s: As a string: �B2UP5#P)�

	s2 := "€£²k1tt"
	// 'x' - номер стартового байта символа, 'y' - сам символ
	for x, y := range s2 {
		fmt.Printf("%#U starts at byte position %d\n", y, x)
		// like: "U+00B2 '²' starts at byte position 5"   - 2b
		//       "U+0074 't' starts at byte position 9"   - 1b
		//       "U+0074 't' starts at byte position 10"  - 1b
	}

	// длина строки - В БАЙТАХ, А НЕ В СИМВОЛАХ
	fmt.Printf("s2 length: %d\n", len(s2)) // 11

	const s3 = "ab12AB"
	fmt.Println("s3:", s3)     // s3: ab12AB
	fmt.Printf("x: % x\n", s3) // x: 61 62 31 32 41 42
	fmt.Printf("s3 length: %d\n", len(s3))
	for i := 0; i < len(s3); i++ {
		fmt.Printf("%x ", s3[i]) // 61 62 31 32 41 42
	}
	fmt.Println()
}
