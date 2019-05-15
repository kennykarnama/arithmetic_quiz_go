package parserutil

import (
	"encoding/csv"
	"os"
)

type CsvParser struct {
	FileLocation string
	Reader       *csv.Reader
}

// Construct new Csv Parser
func NewCsvFile(fileloc string) (*CsvParser, error) {
	csvFile, err := os.Open(fileloc)

	if err != nil {
		return nil, err
	}

	parser := CsvParser{
		FileLocation: fileloc,
		Reader:       csv.NewReader(csvFile),
	}

	return &parser, nil
}

// Query all records inside data
func (parser *CsvParser) QueryRecords() ([][]string, error) {

	records, err := parser.Reader.ReadAll()

	if err != nil {
		return nil, err
	}

	return records, nil
}
