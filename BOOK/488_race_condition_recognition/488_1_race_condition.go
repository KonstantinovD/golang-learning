package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

// go run -race .\488_1_race_condition.go 10
func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("Give me a natural number!")
		return
	}
	numGR, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	var waitGroup sync.WaitGroup
	var i int
	k := make(map[int]int)
	k[1] = 12

	for i = 0; i < numGR; i++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			k[i] = i // WARNING: DATA RACE
		}()
	}

	k[2] = 10 // WARNING: DATA RACE -> because before waitGroup.Wait()
	waitGroup.Wait()
	fmt.Printf("k = %#v\n", k)
}
