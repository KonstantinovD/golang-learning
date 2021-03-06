package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// --- Посмотрим, как можно организовать совместный доступ к данным
// с помощью выделенной для этого горутины.
// --- Несмотря на то что общая память является традиционным способом
// взаимодействия потоков между собой, в стандартный комплект поставки
// Go входят встроенные функции синхронизации, которые позволяют одной
// горутине владеть общей частью данных. Это означает, что другие
// горутины должны отправлять сообщения этой горутине, которая владеет
// общими данными, что предотвращает повреждение данных. Такая горутина
// называется управляющей горутиной.
// --- Согласно терминологии Go, это не коммуникация посредством общей
// памяти, а общая память посредством коммуникации.
// Программа генерирует случайные числа посредством управляющей горутины.

var readValue = make(chan int)  // read random numbers
var writeValue = make(chan int) // write random numbers

// для установки значение общей переменной
func set(newValue int) {
	writeValue <- newValue
}

// для чтения значения сохраненной переменной
func read() int {
	return <-readValue
}

// --- логика программы заключена в реализации функции monitor(), а
// точнее, в операторе select, который управляет работой всей программы.

// --- Когда поступает запрос на чтение, функция read() пытается
// выполнить операцию чтения из канала readValue, который управляется
// функцией monitor(). Результатом операции является текущее значение,
// которое хранится в переменной value.
// --- И наоборот, когда мы хотим изменить сохраненное значение, то
// вызываем функцию set(). Она записывает данные в канал writeValue,
// который также обрабатывается оператором select.
// --- В результате никто не может обратиться к общей переменной
// в обход функции monitor().
func monitor() {
	var value int
	for {
		select {
		case newValue := <-writeValue:
			value = newValue
			fmt.Printf("%d ", value)
		case readValue <- value:
		}
	}
}

func main() {
	var n int
	fmt.Print("Enter number of goroutines: ")
	fmt.Scanf("%d", &n)

	fmt.Printf("Going to create %d random numbers.\n", n)
	rand.Seed(time.Now().Unix())
	go monitor()

	var w sync.WaitGroup

	for r := 0; r < n; r++ {
		w.Add(1)
		go func() {
			defer w.Done()
			set(rand.Intn(10 * n))
		}()
	}
	w.Wait()
	fmt.Printf("\nLast value: %d\n", read())
}

// --- Некоторые предпочитают использовать управляющую горутину вместо
// традиционных методов общего доступа к памяти, потому что считают,
// что реализация с использованием управляющей горутины безопаснее,
// ближе к философии Go и намного чище.
