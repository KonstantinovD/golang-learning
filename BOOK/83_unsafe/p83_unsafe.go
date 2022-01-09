package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var value int64 = 5
	var p1 = &value
	var p2 = (*int32)(unsafe.Pointer(p1))
	// --- здесь используется функция unsafe.Pointer():
	// она позволяет нам на свой страх и риск создать указатель int32
	// с именем p2. Этот указатель ссылается на переменную int64
	// с именем value, доступ к которой осуществляется с помощью
	// указателя p1. Любой указатель Go можно преобразовать
	// в unsafe.Pointer.
	// --- Указатель типа unsafe.Pointer позволяет переопределять систему
	// типов Go. Это позволяет существенно повысить производительность,
	// но может быть опасно, если использовать указатели неправильно или
	// небрежно. Кроме того, так разработчики получают больший
	// контроль над данными.

	fmt.Println("*p1: ", *p1) // 5
	fmt.Println("*p2: ", *p2) // 5
	*p1 = 5434123412312431212
	fmt.Println(value)        // 5434123412312431212
	fmt.Println("*p1: ", *p1) // 5434123412312431212
	fmt.Println("*p2: ", *p2) // -930866580
	*p1 = 54341234            // 54341234
	fmt.Println(value)        // 54341234
	fmt.Println("*p1: ", *p1) // 54341234
	fmt.Println("*p2: ", *p2) // 54341234

	fmt.Println("-----", "test INT max value")
	*p1 = 2147483647          // int max value
	fmt.Println("*p1: ", *p1) // 4294967295
	fmt.Println("*p2: ", *p2) // 4294967295
	*p1 = *p1 + 1             // int max value + 1
	fmt.Println("plus 1")
	fmt.Println("*p1: ", *p1) // 4294967296
	fmt.Println("*p2: ", *p2) // -2147483648  == min int value
	fmt.Println("-----", "test UINT max value")

	*p1 = 4294967295 // uint max value
	var p3 = (*uint32)(unsafe.Pointer(p1))
	var p4 = (*uint32)(unsafe.Pointer(p2))
	fmt.Println("*p1: ", *p1) // 4294967295
	fmt.Println("*p2: ", *p2) // -1
	fmt.Println("*p3: ", *p3) // 4294967295
	fmt.Println("*p4: ", *p4) // 4294967295
	*p1 = *p1 + 1             // int max value + 1
	fmt.Println("plus 1")
	fmt.Println("*p1: ", *p1) // 4294967295
	fmt.Println("*p2: ", *p2) // 4294967295
	fmt.Println("*p3: ", *p3) // 4294967295
	fmt.Println("*p4: ", *p4) // 4294967295
}
