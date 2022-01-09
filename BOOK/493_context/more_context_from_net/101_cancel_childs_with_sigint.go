package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
)

func taskA(ctx context.Context, index int) {
	done := false
	go func() {
		// Keep doing something.
		for i := 0; !done; i++ {
			fmt.Printf("A%d%d\n", index, i)
		}
	}()

	// Wait util context is cancelled.
	<-ctx.Done()
	// Turn on closing flag.
	done = true
}

func taskB(ctx context.Context, index int, someChannel chan<- string) {
loop:
	for i := 0; ; i++ {
		select {
		// Try pushing some message to some channel.
		case someChannel <- fmt.Sprintf("B%d%d\n", index, i):
			continue loop
			// Or when context is cancelled, finish the task.
		case <-ctx.Done():
			break loop
		}
	}
}

func main() {
	// Create application context.
	ctx, cancel := context.WithCancel(context.Background())

	// Fork n of task A with application context.
	for i := 0; i < 10; i++ {
		go taskA(ctx, i)
	}

	aChan := make(chan string)

	// Fork m of task B with application context.
	for i := 0; i < 10; i++ {
		go taskB(ctx, i, aChan)
	}

	// Wait for SIGINT.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig

	// Shutdown. Cancel application context will kill all attached tasks.
	cancel()
}
