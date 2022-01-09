package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// Основной характеристикой протокола TCP является надежность.
// TCP-заголовок каждого пакета включает в себя поля для порта источника
// и порта приемника. Эти два поля в сочетании с IP-адресами источника и
// приемника однозначно идентифицируют каждое TCP-соединение. В этом
// разделе рассмотримTCP-клиен

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide host:port.")
		return
	}

	CONNECT := arguments[1]
	// Функция net.Dial() необходима для подключения к удаленному
	// серверу. Первый параметр функции net.Dial() определяет тип сети,
	// а второй — адрес сервера, который также должен включать в себя
	// номер порта. Допустимые значения для первого параметра —
	// tcp, tcp4 (только для IPv4), tcp6 (только для IPv6), udp,
	// udp4 (только для IPv4), udp6 (только для IPv6), ip,
	// ip4 (только для IPv4), ip6 (только для IPv6), unix (UNIX-сокеты),
	// unixgram и unixpacket
	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text+"\n")

		message, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print("->: " + message)

		if strings.TrimSpace(text) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}
	}
}

// --- Если для тестирования TCPclient.go подключить эту программу к
// TCP-серверу, реализованному с использованием netcat(1), то получим
// такой результат:
/*
$ go run TCPclient.go 8001
dial tcp: address 8001: missing port in address
$ go run TCPclient.go localhost:8001
>> Hello from TCPclient.go!
->: Hi from nc!
>> STOP
->:
TCP client exiting...
*/
