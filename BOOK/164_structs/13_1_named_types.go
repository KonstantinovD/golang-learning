package main

import (
	"fmt"
)

type mile int
type kilometer int
type library []string

func main() {
	// Оператор type позволяет определять именованный тип.
	// Именнованный тип основывается на уже существующем типе
	var distance mile = 5
	fmt.Println(distance)
	distance += 5
	fmt.Println(distance)

	// а зачем создавать новый тип, если он все равно ведет себя как int?
	distanceToEnemy(distance)
	var distance2 kilometer = 5
	// distanceToEnemy(distance2)   // ! ошибка
	// передаваемые данные должны быть явным образом определены
	// в программе как значение типа mile,
	// а не типа int или типа kilometer.
	// С помощью именнованных типов
	// мы придаем типу некоторый дополнительный смысл.
	fmt.Println()
	fmt.Println(distance2)
	fmt.Println()

	//Рассмотрим еще один пример:
	var myLibrary library = library{"Book1", "Book2", "Book3"}
	printBooks(myLibrary)
	// Здесь определен именованный тип library, который по сути
	// представляет срез из строк. Данный тип будет представлять
	// своего рода библиотеку, которая включает книги в виде их строковых
	// названий. С помощью функции printBooks можно вывести все книги из
	// этой библиотеки. При этом эта функция работает именно с типом
	// library, а не с любым срезом из строк.
}

func distanceToEnemy(distance mile) {
	fmt.Println("расстояние до противника:")
	fmt.Println(distance, "миль")
}

func printBooks(lib library) {
	for _, value := range lib {
		fmt.Println(value)
	}
}
