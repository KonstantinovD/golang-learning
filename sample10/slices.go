package main

import (
	"fmt"
)

func main() {
	// Срезы (slice) представляют последовательность элементов одного
	// типа переменной длины. В отличие от массивов длина в срезах
	// не фиксирована и динамически может меняться, то есть
	//можно добавлять новые элементы или удалять уже существующие.
	// The type []T is a slice with elements of type T

	// see diff between arrays and slices:
	// https://www.tugberkugurlu.com/archive/working-with-slices-in-go-golang-understanding-how-append-copy-and-slicing-syntax-work
	// https://blog.golang.org/slices#TOC_2.

	//Срез определяется как и массив (только у него не указывается длина):
	var example []string
	fmt.Println("ln of empty slice: ", len(example))
	var users []string = []string{"Tom", "Alice", "Kate"}
	/* users2 := []string{"Tom", "Alice", "Kate"} */

	fmt.Println(users[2]) // Kate
	fmt.Println("-")
	users[2] = "Katherine"

	for _, value := range users {
		fmt.Println(value)
	}

	// --- С помощью функции make() можно создать срез
	//из нескольких элементов, которые будут иметь значения по умолчанию:
	var val int = 3
	var users2 []string = make([]string, val)
	users2[0] = "Tom - 1"
	users2[1] = "Alice - 2"
	users2[2] = "Bob - 3"
	fmt.Println()

	// --- Для добавления в срез применяется встроенная функция
	// append(slice, value...)
	// Первый параметр функции - срез, в который надо добавить,
	// а второй параметр - обавляемое значение (срез, массив)
	users2 = append(users2, "Kate - 4")
	printSlice(users2)
	fmt.Println()
	users3 := []string{"Pavel - 5", "Roman - 6"}
	// The ... unpacks users3 slice. Without the dots, the code would
	// attempt to append the slice as a whole, which is invalid.
	users2 = append(users2, users3...)
	printSlice(users2)
	fmt.Println()
	users3 = []string{"Karyna - 7", "Egor - 8"}
	users2 = append(users2, users3...)
	printSlice(users2)
	fmt.Println()
	usersArr := [2]string{"Pyotr - 9", "Evgenia - 10"}

	// --- SLICE OPERATOR [a:b]
	// var arrToSlice []string = usersArr[:]
	// convert arr to slice and append it - use slice operator [a:b]
	users2 = append(users2, usersArr[:]...)
	printSlice(users2)
	fmt.Println()
	fmt.Println()

	// --- A slice has both a LENGTH and a CAPACITY.
	// The length - the number of elements slice contains.
	// The capacity - the number of elements in the underlying array,
	// counting from the first element in the slice.

	// The length and capacity can be obtained
	// using the funcs len(s) and cap(s)
	// You can extend a slice's length by re-slicing it,
	// provided it has sufficient capacity.
	s := []int{2, 3, 5, 7, 11, 13}
	printIntSlice(s)
	s = s[:0] // Slice the slice to give it zero length.
	printIntSlice(s)
	s = s[:4] // Extend its length.
	printIntSlice(s)
	s = s[2:]            // Drop its first two values - ! CAPACITY = 4 !
	printIntSlice(s)     // len=2 cap=4
	s = append(s, 9, 13) // len=4 cap=4
	printIntSlice(s)
	s = append(s, 14) // len=5 cap=8
	printIntSlice(s)

	s = append(s, s...) // len=10 cap=16
	printIntSlice(s)

	aval := 18
	sliceWithCapacity := make([]int, 0, aval) // len=0 cap=18
	printIntSlice(sliceWithCapacity)
	sliceWithCapacity = append(sliceWithCapacity, s...) // len=10 cap=18
	printIntSlice(sliceWithCapacity)
	sliceWithCapacity = append(sliceWithCapacity, s...) // len=20 cap=36
	printIntSlice(sliceWithCapacity)
	var conclusion string = "SO, NEW_CAPACITY = CAPACITY * 2"
	fmt.Println(conclusion)
	fmt.Println()

	// --- УДАЛЕНИЕ ЭЛЕМЕНТА (с сохранением порядка)
	// Не эффективно, лучше использовать linked list
	names := []string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}
	//удаляем 4-й элемент
	var n = 3
	names = append(names[:n], names[n+1:]...)
	fmt.Println(names) //["Bob", "Alice", "Kate", "Tom", "Paul", "Mike", "Robert"]

	// --- POINTERS TO SLICES: Method receivers
	// Another way to have a function modify the slice header is to pass a
	// pointer to it:
	fmt.Println("Before: len(slice) =", len(names))
	PtrSubtractOneFromLength(&names)
	fmt.Println("After:  len(slice) =", len(names))
	fmt.Println()

	// --- SLICE COPYING
	newSlice := make([]string, len(names), 2*cap(names))
	copy(newSlice, names)
	newSlice[1] = "defValue"
	printSlice(newSlice)
	fmt.Print(" - copy changed")
	fmt.Println()
	printSlice(names)
	fmt.Print(" - initial slice not changed")
}

func printSlice(users []string) {
	for _, value := range users {
		fmt.Print(value, " ")
	}
}

func printIntSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func PtrSubtractOneFromLength(slicePtr *[]string) {
	slice := *slicePtr
	*slicePtr = slice[0 : len(slice)-1]
}
