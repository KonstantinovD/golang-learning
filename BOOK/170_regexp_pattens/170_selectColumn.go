package main

import (
	"bufio"
	"fmt"
	"github.com/sqweek/dialog"
	"io"
	"os"
	"strings"
)

func main() {
	var column int
	fmt.Print("Enter num of column: ")
	_, err := fmt.Scanf("%d", &column)

	if err != nil {
		fmt.Println("Column value is not an integer:", column)
		os.Exit(1)
	}

	if column < 0 {
		fmt.Println("Invalid Column number!")
		os.Exit(1)
	}

	filename, err := dialog.File().Load()
	if err != nil {
		fmt.Println("Please choose a correct file! ", err)
		os.Exit(1)
	}

	fmt.Println("\t\t", filename)
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("error opening file %s\n", err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file (%s) %s",
				filename, err)
			break
		}
		data := strings.Fields(line) //разбивает строку(delim=space)
		if len(data) >= column {
			fmt.Println(data[column-1])
		}
	}
}
