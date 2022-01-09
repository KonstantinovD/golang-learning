package main

import (
	"fmt"
)

// Метод представляет функцию, связанную с определенным типом.
// Он определяется как функция за исключением того,
// что в определении метода необходимо указать получателя (receiver).
/* func (paramName receiverType) methodName (params) (returnedTypes){
    methodBody
} */

// именованный тип, представляющий срез из строк
type library []string

// Для вывода всех элементов из среза мы можем определить
// следующий метод:
// Используя параметр получателя (здесь - l),
// можно обращаться к получателю
func (l library) printLib() {
	for _, val := range l {
		fmt.Print(val, " ")
	}
	fmt.Println()
}

// Подобным образом мы можем определять методы и для структур:
type person struct {
	name string
	age  int
}

func (p person) print() {
	fmt.Println("Имя:", p.name, " | Возраст:", p.age)
}

// передаем параметры в метод структуры
func (p person) eat(meal string) {
	fmt.Println(p.name, "ест", meal)
}

func main() {
	var lib library = library{"Book1", "Book2", "Book3"}
	lib.printLib()
	fmt.Println()

	var tom = person{name: "Tom", age: 24}
	tom.print()
	tom.eat("котлетки с пюрешкой")
	fmt.Println("------------------")
	fmt.Println()

	// При вызове метода, объект структуры, для которого определен метод,
	// передается в него по значению
	var karl = person{name: "Karl", age: 25}
	fmt.Println("before:", karl.age) // before 24
	tom.updateAgeWrong(33)
	fmt.Println("after:", karl.age) // after 24
	// использовать указатели на структуры:

	fmt.Println("before:", karl.age)
	(&karl).updateAge(35)
	fmt.Println("after:", karl.age)
	//  несмотря на то, что метод updateAge определен для указателя на
	// структуру person, можно вызывать этот метод и для объекта person:
	karl.updateAge(41)
	fmt.Println("after 2:", karl.age)
}

func (p person) updateAgeWrong(newAge int) {
	p.age = newAge
}

func (p *person) updateAge(newAge int) {
	(*p).age = newAge
}

// --- You can only declare a method with a receiver whose type is
// defined in the same package as the method. You cannot declare
// a method with a receiver whose type is defined in another package
// (which includes the built-in types such as int).

// --- There are two reasons to use a pointer receiver.
// - 1) the method can modify the value that its receiver points to.
// - 2) to avoid copying the value on each method call. This can be more
// efficient if the receiver is a large struct, for example.
