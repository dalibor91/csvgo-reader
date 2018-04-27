package csv_reader

import "strings"

type CSVColumn struct {
	value string
}

func (t CSVColumn) ToLower() string {
	return strings.ToLower(t.value)
}


