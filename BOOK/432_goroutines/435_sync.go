package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	fmt.Print("Enter num of goroutines: ")
	fmt.Scanf("%d", &count)
	fmt.Printf("Going to create %d goroutines.\n", count)

	var waitGroup sync.WaitGroup

	fmt.Printf("%#v\n", waitGroup)
	for i := 0; i < count; i++ {
		// Каждый вызов sync.Add() увеличивает счетчик в переменной
		// sync.WaitGroup на единицу
		waitGroup.Add(1)
		// Обратите внимание: очень важно вызвать sync.Add(1) перед
		// оператором go, чтобы предотвратить возможное состояние гонки
		go func(x int) {
			// Когда горутина завершает работу, выполняется функция
			// sync.Done(), которая уменьшает счетчик на единицу
			defer waitGroup.Done()
			fmt.Printf("%d ", x)
		}(i)
	}

	// один из элементов поля state1 в sync.WaitGroup содержит счетчик,
	// который увеличивается и уменьшается в соответствии с вызовами
	// sync.Add() и sync.Done().
	fmt.Printf("%#v\n", waitGroup)
	// Вызов sync.Wait() блокируется до тех пор, пока значение
	// счетчика в соответствующей переменной sync.WaitGroup не
	// станет равным нулю, что дает возможность всем горутинам
	// завершить работу.
	waitGroup.Wait()
	fmt.Println("\nExiting...")
}
