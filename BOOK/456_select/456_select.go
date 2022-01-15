package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ключевое слово select довольно мощное; оно может делать многое в самых
// разных ситуациях. Оператор select в Go похож на switch, но только для
// каналов. На практике это означает, что select позволяет горутине
// дождаться завершения нескольких операций коммуникации. Поэтому
// основное преимущество select состоит в том, что этот оператор дает
// возможность работать с несколькими каналами, используя один блок
// select. В результате, построив соответствующие блоки select, можно
// выполнять неблокирующие операции с каналами.

// Самая большая проблема при использовании нескольких каналов и ключевого слова
// select — это взаимоблокировки. Это означает, что необходимо быть
// особенно осторожными при проектировании и разработке, чтобы избежать
// таких взаимоблокировок.

// также показан пример таймера t := time.NewTimer()
// и time.After()

func gen(min, max int, createNumber chan int, end chan bool) {
	// используется так же, как и time.After()
	notWorkingTimer := time.NewTimer(5 * time.Second)

	// В данном операторе select предусмотрено три варианта
	for {
		// операторы select не требуют ветки default. В качестве «умной»
		// ветки default в этом коде можно рассматривать третью ветку
		// оператора select. Так происходит потому, что time.After()
		// ожидает истечения заданного интервала, после чего передает
		// значение текущего времени по возвращаемому каналу — это
		// разблокирует оператор select, если все остальные каналы по
		// какой-либо причине окажутся заблокированными.
		select {
		case createNumber <- rand.Intn(max-min) + min:
		case <-end:
			// Isn't printed only due to closing main && console
			fmt.Println("Isn't printed")
			close(end)
			// Isn't printed only due to closing main && console
			fmt.Println("Isn't printed - 2")
			// See this printing in 456_select_2.go
			return
		case <-time.After(4 * time.Second):
			fmt.Println("\ntime.After()!")
		case <-notWorkingTimer.C:
			{
				fmt.Println("\nnotWorkingTimer.C!")
			}
		}

		// Оператор select не выполняется последовательно, так как все
		// его каналы проверяются одновременно. Если ни один из каналов,
		// указанных в операторе select, не доступен, то оператор select
		// будет заблокирован, пока не освободится один из каналов. Если
		// доступны сразу несколько каналов оператора select, то Go
		// сделает случайный выбор из набора этих доступных каналов.
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	createNumber := make(chan int)
	end := make(chan bool)

	if len(os.Args) != 2 {
		fmt.Println("Need one integer parameter!")
		return
	}
	n, _ := strconv.Atoi(os.Args[1])
	fmt.Printf("Going to create %d random numbers.\n", n)

	go gen(0, 2*n, createNumber, end)
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", <-createNumber)
	}

	// назначение оператора time.Sleep(5 * time.Second) — предоставить
	// функциям time.After() и gen() достаточно времени, чтобы
	// отработать, вернуть результат и таким образом активизировать
	// соответствующую ветвь оператора select.
	time.Sleep(5 * time.Second)
	fmt.Println("Exiting...")
	end <- true

	//time.Sleep(2 * time.Second)
}

// --- Главным преимуществом оператора select является то, что он может
// подключать, распределять нагрузку и управлять несколькими каналами.
// Поскольку каналы служат для связи между горутинами, select соединяет
// каналы, которые соединяют горутины. Таким образом, оператор select
// является одной из наиболее важных, если не самой важной, частью
// модели конкурентности в Go
