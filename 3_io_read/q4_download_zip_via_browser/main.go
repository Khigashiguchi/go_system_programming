package main

import (
	"net/http"
	"archive/zip"
	"io"
	"strings"
)

func main() {
	http.HandleFunc("/", Index)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment: filename=ascii_sample.zip")

	zipWriter := zip.NewWriter(w)
	defer zipWriter.Close()

	f1, err := zipWriter.Create("f1.txt")
	if err != nil {
		panic(err)
	}
	io.Copy(f1, strings.NewReader("1st file"))
	f2, err := zipWriter.Create("f2.txt")
	if err != nil {
		panic(err)
	}
	io.Copy(f2, strings.NewReader("2nd file"))
}
