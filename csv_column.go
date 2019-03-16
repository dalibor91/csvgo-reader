package csv_reader

import (
	"strings"
	"strconv"
)

type CSVColumn struct {
	value string
	index int
}

func (t CSVColumn) Get() string {
	return t.value
}

func (t CSVColumn) GetIndex() string {
	return t.value
}

func (t *CSVColumn) Set(value string) * CSVColumn {
	t.value = value
	return t
}

func (t *CSVColumn) Apply(callback func(string) string) * CSVColumn {
	return t.Set(callback(t.Get()))
}


func (t CSVColumn) ToLower() string {
	return strings.ToLower(t.Get())
}

func (t CSVColumn) ToInt() (int64, error) {
	return strconv.ParseInt(t.Get(), 0,64)
}

func (t CSVColumn) ToDouble() (float64, error) {
	return strconv.ParseFloat(t.Get(),64)
}

func (t CSVColumn) ToBoolean(compareTo string) (bool, error) {
	return (t.Get() == compareTo), nil
}

