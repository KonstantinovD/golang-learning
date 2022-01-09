package main

import (
	"fmt"
	"time"
)

// -->
// timeout c1
// c2 OK
// -->
func main() {
	c1 := make(chan string)
	go func() {
		time.Sleep(time.Second * 3)
		c1 <- "c1 OK"
	}()

	select {
	case res := <-c1:
		fmt.Println(res) // if timeout isn't reached and c1 is finished
	case <-time.After(time.Second * 1):
		fmt.Println("timeout c1") // if timeout is reached
	}

	c2 := make(chan string)
	go func() {
		time.Sleep(time.Second * 3)
		c2 <- "c2 OK"
	}()

	select {
	case res := <-c2:
		fmt.Println(res) // if timeout isn't reached and c2 is finished
	case <-time.After(time.Second * 4):
		fmt.Println("timeout c2") // if timeout is reached
	}
}
