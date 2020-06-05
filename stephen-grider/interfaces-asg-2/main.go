package main

import (
	"io"
	"os"
)

func main() {
	fileName := os.Args[1]
	file, error := os.Open(fileName)
	if error != nil {
		os.Exit(1)
	}
	io.Copy(os.Stdout, file)
}
