package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

// Reads all the contents of the provided CSV file into a [][]string value.
func readAllCSV(fileName string, path string) ([][]string, error) {
	var file, err = os.Open(fmt.Sprintf("%s%s%s", path, string(os.PathSeparator), fileName))
	if err != nil {
		return nil, err
	}

	var reader = csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	err = file.Close()
	if err != nil {
		return nil, err
	}

	return records, nil
}

// Reads the headers (first row) of the provided CSV into a []string value.
func readHeadersCSV(fileName string, path string) ([]string, error) {
	var file, err = os.Open(fmt.Sprintf("%s%s%s", path, string(os.PathSeparator), fileName))
	if err != nil {
		return nil, err
	}

	var reader = csv.NewReader(file)

	// When Read() is called on a particular reader, it reads a single line, reading the next line each time it is called. Hence, if called only once, it gets only the first row.
	headers, err := reader.Read()
	if err != nil {
		return nil, err
	}

	err = file.Close()
	if err != nil {
		return nil, err
	}

	return headers, nil
}
