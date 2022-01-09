package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// --- Context Timeouts ---
// -- Any application that needs to maintain an SLA (service level
// agreement) for the maximum duration of a request, should use time
// based cancellation.
// For example, consider making an HTTP API call to an external service.
// If the service takes too long, it’s better to fail early and cancel
// the request:

func main() {
	// Create a new context
	// With a deadline of 100 milliseconds
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 200*time.Millisecond)

	// Make a request, that will call the google homepage
	req, _ := http.NewRequest(
		http.MethodGet, "http://google.com", nil)
	// Associate the cancellable context we just created to the request
	req = req.WithContext(ctx)

	// Create a new HTTP client and execute the request
	client := &http.Client{}
	res, err := client.Do(req)
	// If the request failed, log to STDOUT
	if err != nil {
		fmt.Println("Request failed:", err)
		return
	}
	// Print the status code if the request succeeds
	time.Sleep(time.Second)
	fmt.Println("Response received, status code:", res.StatusCode)
}
