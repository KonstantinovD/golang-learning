package main

import (
	"fmt"
	"time"
)

// Сигнальный канал — это канал, который применяется только для передачи
// сигналов. Его можно использовать в тех случаях, когда вы хотите
// проинформировать другую программу о чем-либо. К сигнальным каналам
// не нужно прибегать для передачи данных.

// When doing signaling, it’s good practice to use an empty struct as
// the type of the channel, since they’re pretty useless for anything
// other than signalling — it expresses intent in your code.

// Code can block waiting for something to be sent on the channel:
// <-signal

// func1 : Wait for something to finish
// By blocking on a signal channel, you can wait for a task in another
// goroutine to finish:
func func1(signaling chan struct{}) {
	done := make(chan struct{})

	go func() {
		doLongRunningThing()
		close(done) //send signal that work finished
	}()
	// do some other bits
	printLog("[func1()]: start waiting for doLongRunningThing()")

	<-done // wait for that long running thing to finish

	// do more things
	printLog("[func1()]: finish waiting for doLongRunningThing()")

	close(signaling)
}

// func2 -> Start lots of things at the same time
func func2() {
	start := make(chan struct{})
	for i := 0; i < 100; i++ {
		go func(i int) {
			<-start           // wait for the start channel to be closed
			fmt.Print(i, " ") // do something
		}(i)
	}
	// at this point, all goroutines are ready to go - we just need to
	// tell them to start by closing the start channel
	printLog("[func2()]: ready to start printing")
	close(start)
	printLog("[func2()]: func2 finished") //before any go 'fmt.print'
}

func main() {

	func1chan := make(chan struct{})
	printLog("[main()]: start func1 execution")
	go func1(func1chan)
	<-func1chan // wait for func1 (even without another 'go' goroutine)
	printLog("[main()]: finish func1 execution\n")

	// --- Wait for something to finish
	go func2()
	printLog("[main()]: sleep 2 seconds")
	time.Sleep(2 * time.Second)
	fmt.Println()
	printLog("[main()]: sleep finished\n")

	// --- Stopping things
	stopChan := make(chan struct{})
	printLog("[main()]: start longFunc execution")
	go longFunc(stopChan)
	printLog("[main()]: sleep 3 seconds")
	time.Sleep(3 * time.Second)
	printLog("[main()]: sleep finished")
	close(stopChan)
	printLog("[main()]: finish longFunc execution " +
		"(+ wait 1s for longFunc messages)")
	time.Sleep(1 * time.Second) // wait for longFunc closing messages
}

func doLongRunningThing() {
	time.Sleep(2 * time.Second)
}

func printLog(msg string) {
	fmt.Println(time.Now().Format("2006-01-02 3:4:5"), msg)
}

func longFunc(signaling chan struct{}) {

	longFuncChan := make(chan int)

	go func(out chan<- int) {
		out <- 1
		time.Sleep(time.Millisecond * 20)
		printLog("[inner longFunc]: sleep 1 second")
		time.Sleep(time.Second)
		out <- 2
		time.Sleep(time.Millisecond * 20)
		printLog("[inner longFunc]: sleep 1 second")
		time.Sleep(time.Second)
		out <- 3
		time.Sleep(time.Millisecond * 20)
		printLog("[inner longFunc]: sleep 1 second")
		time.Sleep(time.Second)
		out <- 4
		time.Sleep(time.Millisecond * 20)
		printLog("[inner longFunc]: sleep 1 second")
		time.Sleep(time.Second)
		out <- 5
		time.Sleep(time.Millisecond * 20)
		printLog("[inner longFunc]: finished inner longFunc")
	}(longFuncChan)

	printLog("[longFunc]: start execution")

loop:
	for {
		select {
		case v := <-longFuncChan:
			fmt.Println(v)
		case <-signaling: // triggered when the stop channel is closed
			printLog("[longFunc]: stop signal is received")
			break loop // exit
		case <-time.After(7 * time.Second):
			printLog("[longFunc]: timeout")
			break loop // exit
		}
	}

	printLog("[longFunc]: finish execution normal")
}
