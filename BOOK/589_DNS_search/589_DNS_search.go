package main

// Enter ip address: 127.0.0.1
// Enter ip address: youtube.com

import (
	"fmt"
	"net"
)

// Аббревиатура DNS расшифровывается как Domain Name System — система
// доменных имен — и означает способ преобразования IP-адреса в имя,
// такое как packt.com, и обратно из имени в IP-адрес.
// Логика утилиты DNS.go, которая рассмотрена в этом разделе, довольно
// проста: если заданный аргумент командной строки является корректным
// IP-адресом, то программа обработает его как IP-адрес. В противном
// случае программа будет считать, что ей передано имя хоста, которое
// необходимо преобразовать в один или несколько IP-адресов.

// lookIP() получает в качестве входного аргумента IP-адрес и возвращает
// список имен, которые соответствуют этому IP-адресу, полученный
// с помощью функции net.LookupAddr()
func lookIP(address string) ([]string, error) {
	hosts, err := net.LookupAddr(address)
	if err != nil {
		return nil, err
	}
	return hosts, nil
}

func lookHostname(hostname string) ([]string, error) {
	IPs, err := net.LookupHost(hostname)
	if err != nil {
		return nil, err
	}
	return IPs, nil
}

func main() {
	var input string
	fmt.Print("Enter ip address: ")
	fmt.Scanf("%s", &input)

	// выполняет синтаксический анализ строки как адреса IPv4 или IPv6.
	// Если IP-адрес оказывается недействительным, то возвращает nil.
	IPaddress := net.ParseIP(input)

	if IPaddress == nil {
		IPs, err := lookHostname(input)
		if err == nil {
			for _, singleIP := range IPs {
				fmt.Println("single ip:", singleIP)
			}
		}
	} else {
		hosts, err := lookIP(input)
		if err == nil {
			for _, hostname := range hosts {
				fmt.Println("hostname: ", hostname)
			}
		}
	}
}
