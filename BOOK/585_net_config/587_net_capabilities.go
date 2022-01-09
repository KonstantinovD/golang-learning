package main

import (
	"fmt"
	"net"
)

// Назначение netCapabilities.go — раскрыть возможности всех сетевых
// интерфейсов, обнаруженных в системе

func main() {
	// возвращает все интерфейсы текущего компьютера в виде среза,
	// содержащего элементы типа net.Interface
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, interf := range interfaces {
		fmt.Printf("Name: %v\n", interf.Name)
		fmt.Println("Interface Flags:", interf.Flags.String())
		fmt.Println("Interface MTU:", interf.MTU)
		fmt.Println(
			"Interface Hardware Address:", interf.HardwareAddr)
		fmt.Println()
	}
}
