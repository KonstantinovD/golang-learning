package main

// go run .\503_workers_pool.go 15 5
import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

// --- Пул обработчиков — это множество потоков, предназначенных для
// обработки назначаемых им заданий. Веб-сервер Apache и Go-пакет
// net/http работают приблизительно так: основной процесс принимает все
// входящие запросы, которые затем перенаправляются рабочим процессам
// для обработки. Как только рабочий процесс завершает свою работу, он
// готов к обслуживанию нового клиента.
// --- Однако здесь есть главное различие: пул обработчиков использует
// не потоки, а горутины. Кроме того, потоки обычно не умирают после
// обработки запросов, потому что затраты на завершение потока и
// создание нового слишком высоки, тогда как горутина прекращает
// существовать после завершения работы. Пулы обработчиков в Go
// реализованы с помощью буферизованных каналов, поскольку они позволяют
// ограничить число одновременно выполняемых горутин.

// --- Эта программа обрабатывает целые числа и выводит их квадраты
// посредством отдельной горутины для обслуживания каждого запроса

// Client - структура для назначения уникального идентификатора каждому
//запросу, который мы намерены обработать
type Client struct {
	id      int
	integer int
}

// Client содержит входные данные каждого запроса,

// Data содержит результаты запроса
type Data struct {
	job    Client
	square int
}

var (
	size    = 10
	clients = make(chan Client, size)
	data    = make(chan Data, size)
)

func worker(w *sync.WaitGroup) {
	// ВАЖНО! переменная структуры не может быть nil!!!
	// В отличие от указателя на структуру -> can point to nil
	for c := range clients {
		square := c.integer * c.integer
		output := Data{c, square}
		data <- output
		time.Sleep(time.Second)
	}
	w.Done()
}

// Назначение функции makeWP() состоит в том, чтобы сгенерировать
// необходимое количество горутин worker() для обработки всех запросов
func makeWP(n int) {
	var w sync.WaitGroup
	for i := 0; i < n; i++ {
		w.Add(1)
		go worker(&w)
	}
	w.Wait()
	close(data)
}

// Назначение функции create() — правильно создать все запросы,
// используя тип Client, а затем записать их в канал clients для
// обработки. Обратите внимание, что канал clients читается
// функцией worker().
func create(n int) {
	for i := 0; i < n; i++ {
		// ВАЖНО! переменная структуры не может быть nil!!!
		// В отличие от указателя на структуру -> can point to nil
		c := Client{i, i}
		clients <- c
	}
	// see 442_read_from_closed_chan.go
	close(clients)
	fmt.Print("Client requests are received\n\n")
}

func main() {
	// можно использовать функцию cap(), чтобы определить
	// пропускную способность канала
	fmt.Println("Capacity of clients:", cap(clients))
	fmt.Println("Capacity of data:", cap(data))

	if len(os.Args) != 3 {
		fmt.Println("Need #jobs and #workers!")
		os.Exit(1)
	}

	nJobs, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	nWorkers, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}

	// Если количество обработчиков больше, чем размер
	// буферизованного канала clients, то количество создаваемых
	// горутин будет равно размеру канала clients. Аналогично если
	// количество заданий больше, чем количество обработчиков, то
	// задания будут обрабатываться меньшими группами

	// вызов create() для имитации клиентских запросов для обработки
	go create(nJobs)
	finished := make(chan interface{}) // для блокировки main
	go func() {
		for d := range data {
			fmt.Printf("Client ID: %d\tint: ", d.job.id)
			fmt.Printf("%d\tsquare: %d\n", d.job.integer, d.square)
		}
		finished <- true
	}()
	makeWP(nWorkers)                 // выполняет обработку запросов
	fmt.Printf(": %v\n", <-finished) // блокировка main
}
