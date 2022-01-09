package main

// Enter a domain name: mtsoukalos.eu
// Enter a domain name: golang.com

import (
	"fmt"
	"net"
)

// Один из самых популярных DNS-запросов связан с поиском серверов
// доменных имен (name servers), данные о которых хранятся в NS-записях
// этого домена.

func main() {
	var input string
	fmt.Print("Enter a domain name: ")
	fmt.Scanf("%s", &input)

	domain := input
	NSs, err := net.LookupNS(domain)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, NS := range NSs {
		fmt.Println(NS.Host)
	}
}
