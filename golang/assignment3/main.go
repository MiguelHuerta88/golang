package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// read file to open from args
	f := os.Args[1]

	resp, err := os.Open(f)
	if err != nil {
		fmt.Println("Error: Ran into issues when opening up: ", f)
		os.Exit(1)
	}
	// File type implements reader interface. so we can dump to stdOut
	// or we couldve created our own how we did in the http folder
	io.Copy(os.Stdout, resp)
}
