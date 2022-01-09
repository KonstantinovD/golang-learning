package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// --- Emitting a Cancellation Event ---
// -- If you have an operation that could be cancelled, you will have to
// emit a cancellation event through the context.
// -- This can be done using the WithCancel() function, which returns
// a context object, and a function.
// -- Consider the case of 2 dependent operations. Here, “dependent”
// means if one fails, it doesn’t make sense for the other to complete.
// If we get to know early on that one of the operations failed, we
// would like to cancel all dependent operations.

func operation1(ctx context.Context) error {
	// Let's assume that this operation failed for some reason
	// We use time.Sleep to simulate a resource intensive operation
	time.Sleep(100 * time.Millisecond)
	return errors.New("failed")
}

func operation2(ctx context.Context) {
	// We use a similar pattern to the HTTP server
	// that we saw in the earlier example
	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Println("done")
	case <-ctx.Done():
		fmt.Println("halted operation2")
	}
}

func main() {
	// Create a new context
	ctx := context.Background()
	// Create a new context, with its cancellation function
	// from the original context
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		err := operation1(ctx)
		// If this operation returns an error
		// cancel all operations using this context
		if err != nil {
			cancel()
		}
	}()

	// Run operation2 with the same context we use for operation1
	operation2(ctx)
}
