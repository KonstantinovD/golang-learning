package main

import (
	"fmt"
	"time"
)

// Канал каналов — это особая разновидность переменной канала, которая
// вместо обычных типов переменных работает с другими каналами

var times int

func f1(cc chan chan int, f chan bool) {
	c := make(chan int)
	cc <- c
	defer close(c)

	sum := 0
	select {
	case x := <-c:
		for i := 0; i <= x; i++ {
			sum = sum + i
		}
		c <- sum
		if sum%2 == 0 {
			c <- sum
			time.Sleep(time.Millisecond * 400) //to show delay
			// for 'range' func (line 48)
			c <- sum
		}
	case <-f:
		return
	}
}

func main() {
	var times int
	fmt.Print("Enter times: ")
	fmt.Scanf("%d", &times)

	cc := make(chan chan int)

	for i := 1; i < times+1; i++ {
		f := make(chan bool)
		go f1(cc, f)
		ch := <-cc
		ch <- i
		for sum := range ch {
			// will process every new added number (even with delay)
			// while sleeping continues
			fmt.Print("Sum(", i, ")=", sum, " ")
			// Sum(4)=10 Sum(4)=10  ...[delay 0.4s]...  Sum(4)=10
		}
		fmt.Println()
		time.Sleep(time.Second)
		close(f)
	}
}
