package main

import (
	"os"
	"io"
)

func main() {
	fileName := os.Args[1]
	oldFile, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer oldFile.Close()
	newFile, err := os.Create("new.txt")
	if err != nil {
		panic(err)
	}
	defer newFile.Close()
	io.Copy(newFile, oldFile)
}
