package main

import (
	"fmt"
)

// Один пакет может состоять из нескольких файлов.
// Например, определим два файла: 15_1_factorial.go и 15_2_packages_and_import.go
//
// В файле 15_1_factorial.go определим функцию для подсчета факториала.
// Данный файл принадлежит пакету main.
//
// В файле 1_main.go используем функцию для вычисления факториала
//
// потом билдим все:
// go build -o main.exe 15_1_factorial.go 15_2_packages_and_import.go

func main() {
	fmt.Println(factorial(4)) // 24
	fmt.Println(factorial(5)) // 120

	fmt.Print("Press 'Enter' to continue...")
	fmt.Scanln() // wait for Enter Key
}

// Файлы в разных пакетах - let see go modules
