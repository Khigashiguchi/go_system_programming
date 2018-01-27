package main

import (
	"fmt"
)

// Talker interface
type Talker interface {
	Talk()
}

// Greeter struct
type Greeter struct {
	name string
}

// Talk function
func (g Greeter) Talk() {
	fmt.Printf("Hello, my name is %s\n", g.name)
}

func main() {
	var talker Talker
	talker = &Greeter{"wozozo"}
	talker.Talk()
}
