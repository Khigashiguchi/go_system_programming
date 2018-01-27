package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(file, "Write number %d \n", 3)
	fmt.Fprintf(file, "Write string %s \n", "sample string")
	fmt.Fprintf(file, "Write float %f \n", 0.01)
	file.Close()
}
