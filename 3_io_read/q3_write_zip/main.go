package main

import (
	"os"
	"archive/zip"
	"io"
	"strings"
)

// Answer: http://ascii.jp/elem/000/001/252/1252961/
// Reference: http://twinbird-htn.hatenablog.com/entry/2017/08/06/025738
func main() {
	dest, err := os.Create("sample.zip")
	if err != nil {
		panic(err)
	}

	zipWriter := zip.NewWriter(dest)
	defer zipWriter.Close()

	_, err = zipWriter.Create("sample1.txt")
	if err != nil {
		panic(err)
	}
	a, err := zipWriter.Create("a.txt")
	if err != nil {
		panic(err)
	}
	io.Copy(a, strings.NewReader("プログラムで渡すテキストです"))
}
