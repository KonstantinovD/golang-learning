package main

import (
	"fmt"
	"net"
)

func main() {
	// возвращает все интерфейсы текущего компьютера в виде среза,
	// содержащего элементы типа net.Interface
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, interf := range interfaces {
		fmt.Printf("Interface: %v\n", interf.Name)
		byName, err := net.InterfaceByName(interf.Name)
		if err != nil {
			fmt.Println(err)
		}
		addresses, err := byName.Addrs()
		for k, v := range addresses {
			fmt.Printf("Interface Address #%v: %v\n", k, v.String())
		}
		fmt.Println()
	}
}
