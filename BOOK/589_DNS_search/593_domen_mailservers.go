package main

// Enter a domain name: golang.com
import (
	"fmt"
	"net"
)

// Другой очень популярный DNS-запрос связан с получением MX-записей
// домена. MX-записи указывают на почтовые серверы (mail servers) домена

func main() {
	var input string
	fmt.Print("Enter a domain name: ")
	fmt.Scanf("%s", &input)

	domain := input
	MXs, err := net.LookupMX(domain)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, MX := range MXs {
		fmt.Println(MX.Host)
	}
}
