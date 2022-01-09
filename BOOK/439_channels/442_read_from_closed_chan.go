package main

import (
	"fmt"
)

func main() {
	willClose := make(chan int, 10)
	willClose <- -12
	willClose <- 7
	willClose <- 13

	<-willClose
	<-willClose

	close(willClose)
	read := <-willClose
	fmt.Println(read) // 13
	read = <-willClose
	fmt.Println(read) // 0
	// при чтении данных из закрытого канала сначала возвращаются
	// все значения из канала, а потом нулевое значение соответствующего
	// типа, в данном случае 0 ("", nil, etc)

	// Попытка записи в закрытый канал приведет к панике
	// willClose <- 21

}
