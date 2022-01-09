package main

// Горутина — это минимальная сущность Go, которая может быть выполнена
// конкурентно. Использование слова «минимальная» здесь очень важно,
// поскольку горутины не являются автономными сущностями, такими как
// процессы UNIX, — горутины живут в потоках UNIX, которые, в свою
// очередь, живут в процессах UNIX. Основным преимуществом горутин
// является то, что они чрезвычайно легкие, так что запуск тысяч или
// сотен тысяч горутин на одной машине не является проблемой
// --- показаны два способа создания горутин. Первый из них заключается
// в использовании обычных функций, а второй —
// в использовании анонимных функций. Эти способы эквивалентны

import (
	"fmt"
	"time"
)

func function() {
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
}

func main() {
	// код начинается с выполнения function() в качестве горутины. После
	// этого программа продолжает выполнение, в то время как function()
	// начинает работать в фоновом режиме
	go function()

	// создаем горутину, используя анонимную функцию
	go func() {
		for i := 10; i < 20; i++ {
			fmt.Print(i, " ")
		}
	}()

	time.Sleep(1 * time.Second)
	fmt.Println()
}
