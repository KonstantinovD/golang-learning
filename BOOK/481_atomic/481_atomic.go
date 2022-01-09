package main

import (
	"fmt"
	flag "github.com/spf13/pflag"
	"sync"
	"sync/atomic"
)

// --- Атомарная операция — это операция, которая выполняется за один
// шаг относительно других потоков или в данном случае других горутин.
// Таким образом, атомарная операция не может быть прервана на середине.
// --- В стандартную библиотеку Go входит пакет atomic, который иногда
// позволяет обойтись без использования мьютекса. Однако мьютексы более
// универсальны, чем атомарные операции. С помощью пакета atomic можно
// создавать атомные счетчики, доступные для нескольких горутин, не
// рискуя получить проблемы синхронизации и состояние гонки.

type atomCounter struct {
	val int64
}

func (c *atomCounter) Value() int64 {
	return atomic.LoadInt64(&c.val)
}

func main() {
	minusX := flag.Int("x", 100, "Goroutines")
	minusY := flag.Int("y", 200, "Value")
	flag.Parse()
	X := *minusX
	Y := *minusY

	var waitGroup sync.WaitGroup
	counter := atomCounter{}
	counter2 := atomCounter{}

	for i := 0; i < X; i++ {
		waitGroup.Add(1)
		go func(no int) {
			defer waitGroup.Done()
			for i := 0; i < Y; i++ {
				// Требуемая переменная изменяется с помощью
				// функции atomic.AddInt64().
				atomic.AddInt64(&counter.val, 1)
				counter2.val++
			}
		}(i)
	}

	waitGroup.Wait()
	fmt.Println("atomic incrementation:", counter.Value())
	fmt.Println("non-atomic incrementation:", counter2.Value())
}
