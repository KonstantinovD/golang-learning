package main

import (
	"fmt"
	"strconv"
)

// Структуры - тип данных, для представления каких-либо объектов.
// Структуры содержат набор полей различных атрибутов объекта.
type person struct {
	name string
	age  int
}

func main() {
	// Структура представляет новый тип данных,
	// и мы можем определить переменную данного типа:
	var tom person

	// --- С помощью инициализатора можно передать структуре
	// начальные значения:
	tom = person{"Tom", 23}
	printPerson(tom)

	tom2 := person{}  //свойства структуры получат значения по умолчанию:
	printPerson(tom2) // empty string + 0
	// tom2 := person{"Tom"} - error - pass all or noting
	var tom3 = person{age: 24} // only age
	printPerson(tom3)
	bob := person{name: "Bob", age: 31}
	printPerson(bob)
	fmt.Println()

	// --- Можно создать УКАЗАТЕЛЬ на структуру.
	fmt.Println("--- Pointers to struct ---")
	var pavel *person = &person{name: "Pavel", age: 42}
	var evgenia *person = new(person)
	printPerson(*pavel)
	// Для обращения к полям структуры через указатель можно использовать
	//  операцию разыменования
	fmt.Println((*pavel).name)
	// либо напрямую обращаться по указателю
	fmt.Println(pavel.name)
	printPerson(*evgenia)
	fmt.Println("--- Pointers to struct fields ---")
	// можно определять указатели на отдельные поля структуры:
	karl := person{name: "Karl", age: 19}
	var agePointer *int = &karl.age // указатель на поле tom.age
	*agePointer = 35                // изменяем значение поля
	fmt.Println(karl.age)           //  35
	fmt.Println()

	fmt.Println("--- Inner structures ---")
	var kels = phonePerson{
		name: "Tom",
		age:  24,
		contactInfo: contact{
			email: "kels@gmail.com",
			phone: "+1234567899",
		},
	}
	printPhonePerson(kels)
	kels.contactInfo.email = "super@gmail.com"
	printPhonePerson(kels)
	fmt.Println("-----------------")
	fmt.Println()

	fmt.Println("--- Pointers to structs (can be ptr to same struct type) ---")
	// структура не может иметь поле, которое есть тип этой же структуры
	/* type node struct{
	    value int
	    next node
	} */ // - ERROR
	// Вместо этого поле должно быть указателем на структуру:
	// (see node struct)
	first := node{value: 4}
	second := node{value: 5}
	third := node{value: 6}

	first.next = &second
	second.next = &third

	var current *node = &first

	var sep string = ", "
	for current != nil {
		if current.next == nil {
			sep = ""
		}
		fmt.Print(current.value, sep)
		current = current.next
	}
	fmt.Println()
	fmt.Println("--- Print nodes recursively ---")
	printNodeValue(&first)
}

// Поля одних структур могут представлять другие структуры. Например:
type contact struct {
	email string
	phone string
}

type phonePerson struct {
	name        string
	age         int
	contactInfo contact
}

// Вместо этого поле должно представлять указатель на структуру
type node struct {
	value int
	next  *node
}

// рекурсивный вывод списка
func printNodeValue(n *node) {

	fmt.Println(n.value)
	if n.next != nil {
		printNodeValue(n.next)
	}
}

func printPerson(p person) {
	fmt.Println(p.name, strconv.Itoa(p.age))
}
func printPhonePerson(p phonePerson) {
	fmt.Println(p.name, strconv.Itoa(p.age),
		"|", p.contactInfo.email, "|", p.contactInfo.phone)
}
