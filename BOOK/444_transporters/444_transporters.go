package main

// Конвейер — это виртуальный метод, предназначенный для соединения
// горутин и каналов, так что выходные данные одной горутины становятся
// входными данными для другой горутины, а для передачи данных
// используются каналы
// Одним из преимуществ использования конвейеров является наличие
// постоянного потока данных, так что никакие горутины и каналы не
// должны ожидать, пока завершится все остальное, чтобы можно было
// начать выполнение. Кроме того, мы используем меньше переменных и,
// следовательно, меньше памяти, потому что не приходится сохранять все
// данные в виде переменных. Наконец, использование конвейеров упрощает
// разработку программ и делает их удобнее для поддержки

// --- Задача, выполняемая программой, состоит в том, чтобы генерировать
// случайные числа в заданном диапазоне и останавливаться, когда любое
// число в этой последовательности встретится во второй раз. Прежде чем
// завершить работу, программа выводит на экран сумму всех случайных
// чисел, сгенерированных до того момента, когда впервые появилось
// повторяющееся случайное число

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var CLOSEA = false
var DATA = make(map[int]bool)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func first(min, max int, out chan<- int) {
	for {
		if CLOSEA {
			close(out)
			return
		}
		out <- random(min, max)
	}
}

// в конвейере всем управляет second()
func second(out chan<- int, in <-chan int) {
	// x range chan
	for x := range in {
		fmt.Print(x, " ")
		_, ok := DATA[x] // check if same number was marked in arr
		if ok {
			CLOSEA = true
		} else {
			DATA[x] = true
			out <- x
		}
	}
	fmt.Println()
	close(out)
}

// Функция third() считывает данные из канала, переданного функции
// в качестве аргумента. Когда second() закрывает этот канал, цикл for
// прекращает получать данные и функция выводит результат на экран
func third(in <-chan int) {
	var sum int
	sum = 0
	for x2 := range in {
		sum = sum + x2
	}
	fmt.Printf("The sum of the random numbers is %d\n", sum)
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Need two integer parameters!")
		return
	}
	n1, _ := strconv.Atoi(os.Args[1])
	n2, _ := strconv.Atoi(os.Args[2])

	rand.Seed(time.Now().UnixNano())
	A := make(chan int)
	B := make(chan int)

	// Здесь мы определяем требуемые каналы и выполняем две горутины и
	// одну функцию. Функция third() не позволяет main() немедленно
	// завершиться, поскольку она не выполняется как горутина
	go first(n1, n2, A)
	go second(B, A)
	third(B)

	// BUT! command
	// go run -race .\444_transporters.go 1 10
	// tells that race is possible -> will be fixed in part 10 of book
}
