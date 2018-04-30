package main

import (
	"bufio"
	"strings"
	"fmt"
)

var source = `一行目
2行目
3行目`

func main() {
	//reader := bufio.NewReader(strings.NewReader(source))
	//for {
	//	line, err := reader.ReadString('\n')
	//	fmt.Printf("%#v\n", line)
	//	if err == io.EOF {
	//		break
	//	}
	//}
	scanner := bufio.NewScanner(strings.NewReader(source))
	for scanner.Scan() {
		fmt.Printf("%v\n", scanner.Text())
	}
}