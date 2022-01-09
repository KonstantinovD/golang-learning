package main

import (
	"fmt"
	"unsafe"
)

func main() {
	arr := [...]int{0, 1, -2, 3, 4}
	ptr := &arr[0]
	fmt.Print(*ptr, " ")
	// pointer, которая указывала на целочисленное значение,
	// преобразуется в unsafe.Pointer(), потом — в uintptr
	memAddress := uintptr(unsafe.Pointer(ptr)) + unsafe.Sizeof(arr[0])

	for i := 0; i < len(arr)-1; i++ {
		ptr = (*int)(unsafe.Pointer(memAddress))
		fmt.Print(*ptr, " ")
		memAddress = uintptr(unsafe.Pointer(ptr)) + unsafe.Sizeof(arr[0])
	}

	fmt.Println()
	ptr = (*int)(unsafe.Pointer(memAddress))
	fmt.Print("One more: ", *ptr, " ")
	memAddress = uintptr(unsafe.Pointer(ptr)) + unsafe.Sizeof(arr[0])
	fmt.Println()
}
