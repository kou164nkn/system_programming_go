package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func CreateCsvAndContent(fileName string, records [][]string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	for _, record := range records {
		if err := writer.Write(record); err != nil {
			os.Remove(fileName)
			return err
		}
	}
	writer.Flush()

	if err := writer.Error(); err != nil {
		os.Remove(fileName)
		return err
	}

	return nil
}

func main() {
	fileName := "tennis.csv"

	records := [][]string{
		{"firstName", "lastName", "country"},
		{"John", "Mackenrow", "USA"},
		{"Roger", "Federer", "Switzerland"},
	}

	if err := CreateCsvAndContent(fileName, records); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
