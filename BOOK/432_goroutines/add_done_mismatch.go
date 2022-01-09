package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	fmt.Print("Enter num of goroutines: ")
	fmt.Scanf("%d", &count)
	fmt.Printf("Going to create %d goroutines.\n", count)
	command := selectAction()

	var waitGroup sync.WaitGroup

	if command == 1 {
		// goroutines are asleep - deadlock!
		waitGroup.Add(1)
	}
	fmt.Printf("%#v\n", waitGroup)
	for i := 0; i < count; i++ {
		waitGroup.Add(1)
		go func(x int) {
			defer waitGroup.Done()
			fmt.Printf("%d ", x)
		}(i)
	}

	if command == 2 {
		// negative WaitGroup counter
		waitGroup.Done()
		// or some goroutine cannot be executed
		// e.g: 9 0 1 2 3 4 5 6 7 - no '8' number
	}

	fmt.Printf("%#v\n", waitGroup)
	waitGroup.Wait()
	fmt.Println("\nExiting...")

}

func selectAction() int {
	var command int
	fmt.Print("Enter 1 for extra 'Add', 2 for extra 'Done': ")
	fmt.Scanf("%d", &command)
	return command
}
