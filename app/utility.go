package app

import (
	"encoding/csv"
	"errors"
	"io"
	"strconv"
	"time"
)

type datum struct {
	Input string
}

func (datum datum) isInt() bool {
	if _, err := strconv.ParseInt(datum.Input, 10, 64); err != nil {
		return false
	}
	return true
}

func (datum datum) isFloat() bool {
	if _, err := strconv.ParseFloat(datum.Input, 64); err != nil {
		return false
	}
	return true
}

func (datum datum) isBool() bool {
	if _, err := strconv.ParseBool(datum.Input); err != nil {
		return false
	}
	return true
}

func (datum datum) isDateTime() bool {
	// more formats can be added to support wide array of
	// date time formats
	formatArray := []string{
		"02/01/2006",
		"2006/01/02",
		"02-01-2006",
		"2006-01-02",
		"02-01-2006 15:04",
		"02-01-2006 15:04:05",
		"2006-01-02 15:04",
		"2006-01-02 15:04:05",
		"02/01/2006 15:04",
		"02/01/2006 15:04:05",
		"2006/01/02 15:04",
		"2006/01/02 15:04:05",
	}
	for _, element := range formatArray {
		if _, err := time.Parse(element, datum.Input); err == nil {
			return true
		}
	}
	return false
}

func inferDatumSchema(input string) string {
	uno := datum{
		Input: input,
	}
	if uno.isInt() {
		return "int"
	}
	if uno.isFloat() {
		return "float"
	}
	if uno.isBool() {
		return "bool"
	}
	if uno.isDateTime() {
		return "datetime"
	}
	return "string"
}

func InferDataSchema(fileStream io.Reader) (map[string]string, error) {
	csvReader := csv.NewReader(fileStream)
	data, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}
	if len(data) < 2 {
		return nil, errors.New("invalid csv file")
	}
	schemaMap := make(map[string]string)
	header := data[0]
	row := data[1]
	for ix, _ := range header {
		schemaMap[header[ix]] = inferDatumSchema(row[ix])
	}
	return schemaMap, nil
}
