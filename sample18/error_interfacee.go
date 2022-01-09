package main

import (
	"fmt"
	"time"
)

// The error type is a built-in interface similar to fmt.Stringer
// Functions often return an error value, and calling code
// should handle it by testing whether the error equals nil.

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
