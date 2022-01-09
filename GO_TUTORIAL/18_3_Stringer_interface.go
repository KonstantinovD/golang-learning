package main

import (
	"fmt"
)

/* 1 - Stringer */
// defined by the fmt package
/*
type Stringer interface {
    String() string
}
*/
// A Stringer is a type that can describe itself as a string.
// fmt package (and many others) look for this interface to print values.

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ipaddr IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d",
		ipaddr[0], ipaddr[1], ipaddr[2], ipaddr[3])
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
	fmt.Println()

	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
