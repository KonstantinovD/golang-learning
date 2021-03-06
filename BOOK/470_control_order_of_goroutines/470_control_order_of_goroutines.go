package main

import (
	"fmt"
	"time"
)

// A заблокирована каналом, который хранится в параметре 'a'.
// Как только этот канал будет разблокирован в main(), функция A()
// начнет работать. В конце она закроет канал 'b', тем самым
// разблокировав другую функцию — в данном случае B()
func A(a, b chan struct{}) {
	<-a
	fmt.Println("A()!")
	time.Sleep(time.Second)
	close(b)
}

// Логика функции B() такая же, как и у A(). Эта функция блокируется,
// пока не будет закрыт канал a. Затем она выполняет свою работу и
// закрывает канал b.

func B(a, b chan struct{}) {
	<-a
	fmt.Println("B()!")
	close(b)
}

// C заблокирована и ожидает закрытия канала a, чтобы начать работу
func C(a chan struct{}) {
	<-a
	fmt.Println("C()!")
}

func main() {
	x := make(chan struct{})
	y := make(chan struct{})
	z := make(chan struct{})

	// Многократный вызов функции C() как горутины не вызовет проблем,
	// потому что C() не закрывает никаких каналов. Но если вызвать A()
	// или B() более одного раза, то выведется сообщение об ошибке
	go C(z)
	go A(x, y)
	go C(z)
	go B(y, z)
	go C(z)
	close(x)
	time.Sleep(3 * time.Second)
}
