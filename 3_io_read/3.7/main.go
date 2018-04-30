package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"fmt"
)

func main() {
	// 3.7 io.MultiReader
	//header := bytes.NewBufferString("----- HEADER -----\n")
	//content := bytes.NewBufferString("Example of io.MultiReader\n")
	//footer := bytes.NewBufferString("----- FOOTER -----\n")
	//
	//reader := io.MultiReader(header, content, footer)
	//io.Copy(os.Stdout, reader)

	// 3.7 io.TeeReader
	var buffer bytes.Buffer
	reader := bytes.NewBufferString("Example of io.TeeReader\n")
	teeReader := io.TeeReader(reader, &buffer)
	_, _ = ioutil.ReadAll(teeReader)
	fmt.Println(buffer.String())
}