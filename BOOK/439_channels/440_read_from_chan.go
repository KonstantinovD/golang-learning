package main

import (
	"fmt"
	"time"
)

func writeToChannel2(c chan int, x int) {
	fmt.Println("1", x)
	c <- x
	close(c)
	fmt.Println("2", x)
}

func main() {
	c := make(chan int)
	go writeToChannel2(c, 10)
	time.Sleep(time.Second)
	// мы читаем данные из канала c, используя нотацию <-c
	fmt.Println("Read:", <-c)
	time.Sleep(time.Second)

	_, ok := <-c
	if ok {
		fmt.Println("Channel is open.")
	} else {
		fmt.Println("Channel is closed!")
	}
}
