package main

import (
	"fmt"
	"sync"
	"time"
)

// Основная работа выполняется функциями sync.Lock() и sync.Unlock(),
// первая из которых блокирует, а вторая разблокирует мьютекс sync.Mutex.
// Когда мьютекс блокирован, это означает, что никто другой не может
// заблокировать этот мьютекс, пока он не будет освобожден с помощью
// функции sync.Unlock().

var (
	m  sync.Mutex
	v1 int
)

func change(i int) {
	m.Lock()
	defer m.Unlock() // if mutex not unlocked (e.g. forgotten) -> panic
	time.Sleep(time.Second)
	v1 = v1 + 1
	if v1%10 == 0 {
		v1 = v1 - 10*i
		fmt.Printf(" [%d] ", i)
	}
}

func read() int {
	m.Lock()
	defer m.Unlock()
	return v1
}

func main() {
	var numGOR int
	fmt.Print("Enter number of goroutines: ")
	fmt.Scanf("%d", &numGOR)

	var waitGroup sync.WaitGroup
	fmt.Printf("%d", read())
	for i := 0; i < numGOR; i++ {
		waitGroup.Add(1)
		go func(i int) {
			defer waitGroup.Done()
			change(i)
			fmt.Printf(" -> %d", read())
		}(i)
	}

	waitGroup.Wait()
	fmt.Printf("\n%d", read())
}
