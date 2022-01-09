package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

// пока все функции, читающие мьютекс типа sync.RWMutex, не разблокируют
// этот мьютекс, вы не сможете заблокировать его для записи

var Password = secret{password: "myPassword"}

type secret struct {
	RWM      sync.RWMutex
	M        sync.Mutex
	password string
}

// Change изменяет общую переменную. Следовательно, в ней нужно
// использовать монопольную блокировку, для чего применяются функции
// Lock() и Unlock()
func Change(c *secret, pass string) {
	c.RWM.Lock()
	defer c.RWM.Unlock()
	fmt.Println("LChange")
	time.Sleep(10 * time.Second)
	c.password = pass

}

func show(c *secret) string {
	c.RWM.RLock()
	defer c.RWM.RUnlock()
	fmt.Println("show")
	time.Sleep(3 * time.Second)
	return c.password

}

func showWithLock(c *secret) string {
	c.RWM.Lock()
	defer c.RWM.Unlock()
	fmt.Println("showWithLock")
	time.Sleep(3 * time.Second)
	return c.password

}

func main() {
	var showFunction = func(c *secret) string { return "" }
	var operat int
	fmt.Print("input '1' for RWMutex, '2' for Mutex: ")
	fmt.Scanf("%d", &operat)

	switch operat {
	case 1:
		fmt.Println("Using sync.RWMutex")
		showFunction = show
	case 2:
		fmt.Println("Using sync.Mutex")
		showFunction = showWithLock
	default:
		fmt.Println("Incorrect input!")
		os.Exit(-1)
	}
	fmt.Println("Pass:", showFunction(&Password))

	var waitGroup sync.WaitGroup
	for i := 0; i < 15; i++ {
		waitGroup.Add(1)
		go func(i int) {
			defer waitGroup.Done()
			fmt.Println(strconv.Itoa(i), " Go Pass:", showFunction(&Password))
		}(i)
	}

	go func() {
		waitGroup.Add(1)
		defer waitGroup.Done()
		Change(&Password, "123456")
	}()

	waitGroup.Wait()
	fmt.Println("Pass:", showFunction(&Password))
}
