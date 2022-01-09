package main

import "fmt"

func f1(c chan int, x int) {
	fmt.Println(x)
	c <- x
}
func f2(c chan<- int, x int) {
	fmt.Println(x)
	c <- x
	// <- c  - ERROR
}

// Go позволяет указать направление канала при применении его в качестве
// аргумента. Эти два типа каналов называются однонаправленными каналами
func f3(out chan<- int64, in <-chan int64) {
	k := <-in
	out <- k
}

func main() {
	out := make(chan int64, 3)
	in := make(chan int64, 3)

	in <- 112
	f3(out, in)
	fmt.Println(<-out) // 112
}
