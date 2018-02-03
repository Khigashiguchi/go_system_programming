package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	records := [][]string{
		{"record1", "record2", "record3"},
		{"record4", "record5"},
		{"record6", "record6", "record7"},
		{"record8", "record9", "record10"},
	}
	writer := csv.NewWriter(os.Stdout)
	for _, record := range records {
		if err := writer.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}
	writer.Flush()
	if err := writer.Error(); err != nil {
		log.Fatal(err)
	}
}
