package main

import (
	"fmt"
	"runtime"
)

func printStats(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)
	fmt.Println("mem.Alloc: ", mem.Alloc)
	fmt.Println("mem.TotalAlloc: ", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc: ", mem.HeapAlloc)
	fmt.Println("mem.NumGC: ", mem.NumGC)
	fmt.Println("-------")
}

func main() {
	var mem runtime.MemStats
	printStats(mem)

	// Цикл for создает много больших срезов Go, чтобы под них
	// выделялись большие объемы памяти и запускался сборщик мусора
	for i := 0; i < 10; i++ {
		s := make([]byte, 50000000)
		if s == nil {
			fmt.Println("Operation failed!")
		}
	}

	printStats(mem)
	/*
		mem.Alloc:  50109952
		mem.TotalAlloc:  500159328
		mem.HeapAlloc:  50109952
		mem.NumGC:  9
	*/

	// вызываем дополнительное выделение памяти для срезов Go:
	for i := 0; i < 10; i++ {
		s := make([]byte, 100000000)
		if s == nil {
			fmt.Println("Operation 2 failed!")
		}
	}

	printStats(mem)
	/*
		mem.Alloc:  100114024
		mem.TotalAlloc:  1500239600
		mem.HeapAlloc:  100114024
		mem.NumGC:  19
	*/
}
