package main

import (
	"fmt"
	"leaning/BOOK/285_go_packages/01_message"
)

func main() {
	// В Go действует простое правило, согласно которому функции,
	// переменные, типы и т. д., имена которых начинаются с заглавной
	// буквы, являются открытыми. А функции, переменные, типы и т. д.,
	// имена которых начинаются со строчной буквы, являются закрытыми.
	// Именно поэтому функция fmt.Println() называется Println(),
	// а не println(). Однако это правило не распространяется на имена
	// пакетов, которые могут начинаться как с прописных, так и со
	// строчных букв

	msg := message.OpenMessage()
	fmt.Println(msg)
	fmt.Println(msg.Value)
	fmt.Println(msg.Sevrt)

	level := message.Level(33)
	fmt.Println("no restrictions for 'enums', level =", level)
}
