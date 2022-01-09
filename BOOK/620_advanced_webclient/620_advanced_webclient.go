package main

// go run .\620_advanced_webclient.go http://www.mtsoukalos.eu
// go run .\620_advanced_webclient.go http://www.google.com

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s URL\n", filepath.Base(os.Args[0]))
		return
	}
	URL, err := url.Parse(os.Args[1])
	if err != nil {
		fmt.Println("Error in parsing:", err)
		return
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	// возвращает объект http.Request с указанием метода, URL-адреса и,
	// возможно, тела запроса
	request, err := http.NewRequest(
		"GET", URL.String(), nil)
	if err != nil {
		fmt.Println("Get:", err)
		return
	}
	// отправляет http.Request через http.Client и получает Response
	httpData, err := client.Do(request)
	if err != nil {
		fmt.Println("Error in Do():", err)
		return
	}

	fmt.Println("Status code:", httpData.Status)
	header, _ := httputil.DumpResponse(httpData, false)
	fmt.Print(string(header))

	contentType := httpData.Header.Get("Content-Type")
	characterSet := strings.SplitAfter(contentType, "charset=")
	if len(characterSet) > 1 {
		fmt.Println("Character Set:", characterSet[1])
	}

	if httpData.ContentLength == -1 {
		fmt.Println("ContentLength is unknown!")
	} else {
		fmt.Println("ContentLength:", httpData.ContentLength)
	}

	length := 0
	var buffer [1024]byte
	r := httpData.Body
	for {
		n, err := r.Read(buffer[0:])
		if err != nil {
			fmt.Println(err)
			break
		}
		length = length + n
	}
	fmt.Println("Calculated response data length:", length)
}
