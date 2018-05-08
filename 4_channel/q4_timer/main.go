package main

import (
	"time"
	"fmt"
)

func main() {
	fmt.Println("start timer ...")
	select {
	case <-time.After(10 * time.Second):
		fmt.Println("time out!")
	}
	fmt.Println("end timer")

}
