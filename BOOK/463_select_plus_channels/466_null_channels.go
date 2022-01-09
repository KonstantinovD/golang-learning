package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Нулевые каналы - это особый вид каналов, который всегда блокирован.

func add(in chan int) {
	sum := 0
	t := time.NewTimer(time.Second)

	for {
		// Оператор <-t.C блокирует канал C таймера t на время,
		// указанное в вызове time.NewTimer(). По истечении времени
		// таймер отправляет значение в канал t.C, чем инициирует
		// выполнение соответствующей ветви оператора select
		select {
		case input := <-in:
			sum = sum + input
		case <-t.C: // timer channel
			in = nil
			fmt.Println(sum)
		}
	}
}

// Цель функции send() — генерировать случайные числа и отправлять их в
// канал до тех пор, пока канал открыт.
func send(c chan int) {
	for {
		c <- rand.Intn(10)
	}
}

func main() {
	c := make(chan int)
	go add(c)
	go send(c)

	// sleep, чтобы у двух горутин хватило времени для выполнения.
	time.Sleep(3 * time.Second)
}
