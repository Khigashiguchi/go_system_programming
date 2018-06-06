package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-type", "application/json")
	source := map[string]string{
		"Hello": "World",
	}
	// json encoding
	jsonBytes, _ := json.Marshal(source)
	fmt.Println(string(jsonBytes))

	// gzip and write os.Stdout
	gz := gzip.NewWriter(w)
	defer gz.Close()
	gz.Header.Name = "test.json"
	writer := io.MultiWriter(gz, os.Stdout)
	io.WriteString(writer, string(jsonBytes))
	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(writer)
	gz.Flush()
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
