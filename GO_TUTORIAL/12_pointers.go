package main

import (
	"fmt"
)

func main() {
	// Этому указателю можно присвоить адрес переменной типа int.
	// Для получения адреса применяется операция &, после которой
	// указывается имя переменной (&x).
	var x int = 4              // определяем переменную
	var p *int                 // определяем указатель
	p = &x                     // указатель получает адрес переменной
	fmt.Println("Address:", p) //значение указателя - адрес переменной x
	fmt.Println("Value:", *p)  // значение переменной x
	fmt.Println("change value by pointer")
	// из указателя можем менять значение по адресу,
	// который в нем хранится:
	*p = 25
	fmt.Println("Address:", p) // значение указателя - адрес переменной x
	fmt.Println("Value:", *p)  // значение переменной x
	fmt.Println()
	// Для определения указателей можно использовать сокращенную форму:
	f := 2.3
	pf := &f
	fmt.Println("Address:", pf)
	fmt.Println("Value:", *pf)
	fmt.Println()

	// Если указателю не присвоен адрес какого-либо объекта, то такой
	// указатель по умолчанию имеет значение nil. Если мы попробуем
	// получить значение по такому пустому указателю, то мы столкнемся
	// с ошибкой. Поэтому при работе с указателями
	// стоит проверять значение на nil:
	var pff *float64
	if pff != nil {
		// or panic: runtime error:
		// invalid memory address or nil pointer dereference
		fmt.Println("Value:", *pff)
	}
	fmt.Println()

	/* --- NEW & MAKE difference*/
	// - first difference is type: make(T, ...) always returns type T,
	// whereas new(T, ...) always returns type *T.
	// - second: new works for all types, and dynamically allocates space
	// for a variable of that type, initialized to the zero value of that
	// type, and returns a pointer to it.

	//One way to think of it is that
	result := new(string)
	fmt.Println("Address:", result)
	fmt.Println("Value:", *result)
	//is always equivalent to
	var temp string
	result = &temp
	fmt.Println("Address:", result)
	fmt.Println("Value:", *result)
	fmt.Println()

	// you can do the same thing that new does by
	// defining a variable of the given type, NOT INITIALIZING IT,
	// and then taking a pointer to it.

	// --- make works as a kind of "constructor" for certain built-in
	// types (slice, map, or channel).

	// --- If new was removed in favour make, how would you construct
	// a pointer to an initialised value?
	var x1 *int // nil
	var x2 = new(int)
	fmt.Println(x1, x2)
	fmt.Println("x2 points to value =", *x2)
	fmt.Println()
	// x1 и x2 имеют один и тот же тип, * int, x2 указывает на
	// инициализированную память и может быть безопасно разыменован,
	// то же самое не верно для x1 -
	// он не может быть разыменован прямо сейчас

	// --- в свою очередь, new не умеет работать
	//со slice, map, or channel
	sl1 := new([]string)
	fmt.Println(sl1 == nil)  // false
	fmt.Println(*sl1 == nil) // true
	sl2 := make([]string, 0)
	fmt.Println(sl2 == nil)                              // false
	fmt.Println("sl2 size=", len(sl2), "cap=", cap(sl2)) // 0 0
	fmt.Println("-----------")

	m := new(map[string]int)
	fmt.Println(len(*m))   // 0
	fmt.Println(m == nil)  // false
	fmt.Println(*m == nil) // true
	fmt.Println("-----------")

	c := new(chan int)
	fmt.Println(c == nil)  // false
	fmt.Println(*c == nil) // true
	fmt.Println("-----------")

}
