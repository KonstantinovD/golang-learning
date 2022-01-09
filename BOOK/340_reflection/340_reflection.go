package main

import (
	"fmt"
	"reflect"
)

// Рефлексия — это дополнительная функциональность Go, которая позволяет
// динамически, в процессе выполнения программы, узнавать тип
// произвольного объекта, а также получать информацию о его структуре.
// Для работы с рефлексией в Go создан пакет reflect.
// Зачем нужна рефлексия и когда ее стоит использовать?
// Рефлексия необходима, например, для реализации таких пакетов, как
// fmt, text/template и html/template. В пакете fmt рефлексия избавляет
// нас от необходимости обрабатывать каждый из существующих типов данных
// по отдельности.
// Рефлексия будет полезной в тех случаях, когда вы хотите обеспечить
// максимально возможную универсальность или когда хотите гарантировать
// возможность работы с типами данных, которые не существуют на момент
// написания кода, но могут появиться в будущем. Кроме того, рефлексия
// удобна при работе со значениями типов, которые не соответствуют
// общему интерфейсу

type a struct {
	X int
	Y float64
	Z string
}

type b struct {
	F int
	G int
	H string
	I float64
}

func printReflect(r reflect.Value) {
	iType := r.Type()
	fmt.Printf("i Type: %s\n", iType)
	fmt.Printf("The %d fields of %s are:\n", r.NumField(), iType)
	for i := 0; i < r.NumField(); i++ {
		fmt.Printf("Field name: %s ", iType.Field(i).Name)
		fmt.Printf("with type: %s ", r.Field(i).Type())
		fmt.Printf("and value %v\n", r.Field(i).Interface())
	}
}

func main() {
	x := 100
	xRefl := reflect.ValueOf(&x).Elem()
	xType := xRefl.Type()
	fmt.Printf("The type of x is %s.\n", xType)
	fmt.Println()

	A := a{100, 200.12, "Struct a"}
	B := b{1, 2, "Struct b", -1.2}
	var r reflect.Value

	r = reflect.ValueOf(&A).Elem()
	printReflect(r)
	fmt.Println()
	r = reflect.ValueOf(&B).Elem()
	printReflect(r)
}
