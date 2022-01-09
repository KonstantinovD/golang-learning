package main

import (
	"fmt"
	"time"
)

type SampleData struct {
	name string
	val  int
}

var (
	signal = make(chan struct{})
)

func processNullPointers(ch <-chan *SampleData) {
	for {
		select {
		case ptr := <-ch:
			if ptr == nil {
				fmt.Println("cannot proceed ptr to nil")
			} else {
				fmt.Printf("%s -> %d\n", ptr.name, ptr.val)
			}
		case <-time.After(time.Second):
			close(signal)
			return
		}
	}
}

func createData(out chan<- *SampleData) {
	c := &SampleData{name: "pwdt", val: 11}
	out <- c
	time.Sleep(time.Millisecond * 250)
	c = nil
	out <- c
	time.Sleep(time.Millisecond * 250)
	c = &SampleData{name: "ktgz", val: 13}
	out <- c
	time.Sleep(time.Millisecond * 250)
	c = nil
	out <- c
	time.Sleep(time.Millisecond * 250)
	c = &SampleData{name: "asd", val: 15}
	out <- c
	fmt.Println("close")
	//close(out) - if that will be uncommented ->
	// L20 "case ptr := <-ch" will repeat infinitely
}

func main() {
	dataChan := make(chan *SampleData)
	go createData(dataChan)
	go processNullPointers(dataChan)
	<-signal
	fmt.Println("Exiting...")
}
