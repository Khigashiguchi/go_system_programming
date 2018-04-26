package main

import (
	"os"
	"io"
)

func main() {
	file, err := os.Open("main.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.Copy(os.Stdout, file)
}
