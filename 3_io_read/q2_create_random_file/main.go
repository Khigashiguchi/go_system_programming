package main

import (
	"os"
	"crypto/rand"
)

func main() {
	c := 1024
	b := make([]byte, c)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}

	newFile, err := os.Create("random.txt")
	if err != nil {
		panic(err)
	}
	defer newFile.Close()
	_, err = newFile.Write(b)
	if err != nil {
		panic(err)
	}
}
