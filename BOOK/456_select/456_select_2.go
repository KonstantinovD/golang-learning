package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// test Inner goroutine
// "Isn't printed" && "Isn't printed 2" will be printed
// due to ending of inner goroutine, not maim goroutine

func gen2(min, max int, createNumber chan int, end chan bool) {
	// В данном операторе select предусмотрено три варианта
	for {
		select {
		case createNumber <- rand.Intn(max-min) + min:
		case <-end:
			// printed due to closing inner goroutine
			fmt.Println("Isn't printed")
			close(end)
			// printed due to closing inner goroutine
			fmt.Println("Isn't printed - 2")
			return
		case <-time.After(4 * time.Second):
			fmt.Println("\ntime.After()!")
		}
	}
}

func innerGoroutine() {
	createNumber := make(chan int)
	end := make(chan bool)

	if len(os.Args) != 2 {
		fmt.Println("Need one integer parameter!")
		return
	}
	n, _ := strconv.Atoi(os.Args[1])
	fmt.Printf("Going to create %d random numbers.\n", n)

	go gen2(0, 2*n, createNumber, end)
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", <-createNumber)
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Exiting...")
	end <- true
}

func main() {
	rand.Seed(time.Now().Unix())
	innerGoroutine()
	time.Sleep(7 * time.Second)
}
