package main

import (
	"fmt"
	"sync"
	"time"
)

func timeout(w *sync.WaitGroup, t time.Duration) bool {
	temp := make(chan int)
	go func() {
		defer close(temp)
		time.Sleep(5 * time.Second)
		w.Wait() // заставит анонимную функцию бесконечно ожидать
		// соответствующую функцию sync.Done(), чтобы завершить работу
	}()

	select {
	case <-temp:
		return false
	case <-time.After(t):
		return true
	}
}

func main() {
	var t int
	fmt.Print("Enter a time duration [more than 5s]: ")
	fmt.Scanf("%d", &t)

	var w sync.WaitGroup
	w.Add(1)

	duration := time.Duration(int32(t)) * time.Second
	fmt.Printf("Timeout period is %s\n", duration)

	if timeout(&w, duration) {
		fmt.Println("Timed out!")
	} else {
		fmt.Println("OK!")
	}

	w.Done() // теперь анонимная функция не будет бесконечно ожидать
	// а будет ждать всего 5 сек
	if timeout(&w, duration) {
		fmt.Println("Timed out!")
	} else {
		fmt.Println("OK!")
	}
}
