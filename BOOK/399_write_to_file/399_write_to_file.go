package main

// files will be created in the root directory of this project
// (near the 'go.mod' file)

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	s := []byte("Data to write\n")
	f1, err := os.Create("f1.txt")
	if err != nil {
		fmt.Println("Cannot create file", err)
		return
	}
	defer f1.Close()
	// функция fmt.Fprintf() позволяет записывать данные в ваши
	// собственные файлы журнала, используя нужный формат
	fmt.Fprintf(f1, string(s))

	f2, err := os.Create("f2.txt")
	if err != nil {
		fmt.Println("Cannot create file", err)
		return
	}
	defer f2.Close()
	// функция f2.WriteString()
	n, err := f2.WriteString(string(s))
	fmt.Printf("wrote %d bytes\n", n)

	f3, err := os.Create("f3.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	// функция bufio.NewWriter() открывает файл для записи,
	// а bufio.WriteString() записывает данные.
	w := bufio.NewWriter(f3)
	n, err = w.WriteString(string(s))
	fmt.Printf("wrote %d bytes\n", n)
	w.Flush()

	// В этом методе для записи данных в файл использована функция
	// io.WriteString().
	f4, err := os.Create("f4.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	n, err = io.WriteString(f4, string(s))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("wrote %d bytes\n", n)

	// Этот метод требует только одного вызова функции
	// ioutil.WriteFile() для записи данных, и не требует
	// использования os.Create().
	f5 := "f5.txt"
	// 0644 -> A FileMode represents a file's mode and permission bits.
	// The bits have the same definition on all systems, so that
	// information about files can be moved from one system
	// to another portably. Not all bits apply to all systems.
	err = ioutil.WriteFile(f5, s, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}
