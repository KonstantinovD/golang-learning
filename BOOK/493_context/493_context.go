package main

import (
	"context"
	"fmt"
	"time"
)

// ABOUT: https://www.sohamkamani.com/golang/context-cancellation-and-values/

// Несколько странный пример, но и ладно

// Главное назначение пакета context — определение типа Context и
// поддержка аннулирования. Бывают случаи, когда по какой-либо причине
// завершить процесс (e.g. timeout). Однако есть возможность добавить
// дополнительную информацию о ваших решениях об остановке.
// Именно это позволяет сделать пакет context

// --- Context — это интерфейс, имеющий четыре метода:
// Deadline(), Done(), Err() и Value().
// --- Функции context.WithCancel(), context.WithDeadline() и
// context.WithTimeout() возвращают производный (дочерний) объект
// Context и функцию CancelFunc.
// --- Вызов функции CancelFunc() удаляет ссылку родителя на дочерний
// объект и останавливает все связанные с ним таймеры. После этого
// сборщик мусора Go может свободно собрать дочерние горутины, у которых
// больше нет связанных с ними родительских горутин.
// --- Чтобы сборка мусора работала правильно, родительская горутина
// должна хранить ссылку на все дочерние горутины. Если дочерняя
// горутина завершится без ведома родительской, начинается утечка памяти,
// которая продолжается до тех пор, пока родительская горутина также не
// будет аннулирована

// context.WithCancel()
func f1(t int) {
	c1 := context.Background() // инициализация пустого объекта Context

	// WithCancel() использует существующий объект Context и создает
	// его потомка с возможностью аннулирования. Также создает канал
	// Done, который можно закрыть либо при вызове функции cancel(),
	// либо когда закрывается канал Done родительского контекста.
	c1, cancel := context.WithCancel(c1)
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel() // вызовется если t > 4
	}()

	select {
	case <-c1.Done():
		fmt.Println("t > 4, f1():", c1.Err()) // вызовется если t > 4
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("t <= 4, f1():", r) // вызовется если t <= 4
	}
	return
}

// context.WithTimeout()
func f2(t int) {
	c2 := context.Background()

	// По окончании времени ожидания автоматически
	// вызывается функция cancel().
	c2, cancel := context.WithTimeout(
		c2, time.Duration(t)*time.Second)

	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <-c2.Done():
		fmt.Println("f2():", c2.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("f2():", r)
	}
	return
}

// context.WithDeadline()
func f3(t int) {
	c3 := context.Background()

	// По истечении 2 секунды автоматически вызывается
	// функция cancel()
	deadline := time.Now().Add(time.Duration(t) * time.Second)

	c3, cancel := context.WithDeadline(c3, deadline)
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <-c3.Done():
		fmt.Println("f3():", c3.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("f3():", r)
	}
	return
}

func main() {
	var delay int
	fmt.Print("Enter a delay (in seconds): ")
	_, err := fmt.Scanf("%d", &delay)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Delay:", delay)

	f1(delay)
	f2(delay)
	f3(delay)
}
