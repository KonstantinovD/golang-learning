package main

// go run .\623_http_connections_delays.go http://localhost:8001
// and need to run 'go run .\slowWWW.go' in separate terminal

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// Рассмотрим, как разрывать сетевые соединения, выполнение которых
// занимает слишком много времени. Подобный способ рассмотрен в главе 10
// [497_context_more_complex.go & 497_request_with_context.go]

var timeout = 5 * time.Duration(time.Second)

func Timeout(network, host string) (net.Conn, error) {
	conn, err := net.DialTimeout(network, host, timeout)
	if err != nil {
		return nil, err
	}
	// SetDeadline() используется в пакете net для того, чтобы задать
	// максимальную длительность операций чтения и записи для данного
	// сетевого соединения. Специфика работы функции SetDeadline()
	// такова, что ее необходимо вызвать перед любой операцией чтения
	// или записи. Следует помнить, что максимальная длительность
	// операций необходима в Go для реализации принудительного
	// прерывания операций. Поэтому не нужно переопределять период
	// ожидания каждый раз, когда приложение получает или отправляет
	// данные
	conn.SetDeadline(time.Now().Add(timeout))
	return conn, nil
}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Usage: %s URL TIMEOUT\n", filepath.Base(os.Args[0]))
		return
	}

	if len(os.Args) == 3 {
		temp, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Using Default Timeout!")
		} else {
			timeout = time.Duration(time.Duration(temp) * time.Second)
		}
	}

	URL := os.Args[1]
	//t := http.Transport{DialContext: (&net.Dialer{
	//	Timeout: timeout,
	//	KeepAlive: 2 * timeout,
	//}).DialContext}
	t := http.Transport{Dial: Timeout}

	client := http.Client{
		Transport: &t,
	}
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
