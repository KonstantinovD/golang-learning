package main

import (
	"fmt"
	"time"
)

func main() {
	var count int
	fmt.Print("Enter num of goroutines: ")
	fmt.Scanf("%d", &count)
	fmt.Printf("Going to create %d goroutines.\n", count)

	for i := 0; i < count; i++ {
		go func(x int) {
			fmt.Printf("%d ", x)
		}(i)
	}

	// Цель оператора time.Sleep() — дать горутинам достаточно времени,
	// чтобы они успели завершить свою работу и их результаты можно
	// было увидеть на экране
	time.Sleep(time.Second)
	fmt.Println("\nExiting...")
}
