package main

import (
	"io"
	"os"
)

func main() {
	mSting := ""
	args := os.Args
	if len(args) == 1 {
		mSting = "Please give me one argument!"
	} else {
		mSting = args[1]
	}
	io.WriteString(os.Stdout, mSting)
	io.WriteString(os.Stdout, "\n")
}
