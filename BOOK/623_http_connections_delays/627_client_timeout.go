package main

// go run .\627_client_timeout.go http://localhost:8001
// and need to run 'go run .\slowWWW.go' in separate terminal
import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

var timeout2 = time.Duration(time.Second) * 5

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please provide a URL")
		return
	}
	if len(os.Args) == 3 {
		temp, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Using Default Timeout!")
		} else {
			timeout2 = time.Duration(time.Duration(temp) * time.Second)
		}
	}
	URL := os.Args[1]
	// Timeout переменной http.Client определяется период ожидания
	client := http.Client{
		Timeout: timeout2,
	}
	client.Get(URL)

	data, err := client.Get(URL)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		defer data.Body.Close()
		_, err := io.Copy(os.Stdout, data.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
