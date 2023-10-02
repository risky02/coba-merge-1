package main

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
	"sync"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err.Error())
	}
}

func main() {
	// go run main.go input.csv output.csv
	// cek os.args

	// open input file
	inputFilePath := os.Args[1]
	inputFile, err := os.Open(inputFilePath)
	failOnError(err, "failed to open input file")

	reader := csv.NewReader(inputFile)
	records, err := reader.ReadAll()
	failOnError(err, "failed to read all records")

	wg := sync.WaitGroup{}
	updatedRecords := make(chan []string)

	for _, record := range records[1:] {
		wg.Add(1)
		go func(record []string) {
			defer wg.Done()
			updatedRecords <- []string{
				strings.ToUpper(record[0]),
				record[1],
				"Mr." + record[2],
			}
		}(record)
	}

	go func() {
		wg.Wait()
		close(updatedRecords)
	}()

	// create output file
	outputFilePath := os.Args[2]
	generateOutput(outputFilePath, updatedRecords)
}