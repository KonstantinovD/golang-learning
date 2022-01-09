package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

// --- Listening For Cancellation ---
// -- The Context type provides a Done() method. This returns a channel
// that receives an empty struct{} type every time the context receives
// a cancellation event.
// -- So, to listen for a cancellation event, we need to wait
// on <- ctx.Done().
// -- For example, lets consider an HTTP server that takes two seconds
// to process an event. If the request gets cancelled before that,
// we want to return immediately:

func main() {
	// Create an HTTP server that listens on port 8000
	http.ListenAndServe(":8000",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			fmt.Fprint(os.Stdout, "processing request\n")
			// We use `select` to execute a piece of code depending on
			// which channel receives a message first
			select {
			case <-time.After(5 * time.Second):
				w.Write([]byte("request processed"))
			case <-ctx.Done():
				fmt.Fprint(os.Stderr, "request cancelled\n")
			}

		}))
}
