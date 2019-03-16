package csv_reader

import (
	"errors"
	"strings"
)

type CSVRow struct {
	row []CSVColumn
}

func newCSVRow(row []string) (* CSVRow) {
	k := make([]CSVColumn, 0)
	for i := range(row) {
		k = append(k, CSVColumn{value: row[i], index: i})
	}
	return NewCSVRow(k)
}

func NewCSVRow(row []CSVColumn) (* CSVRow) {
	return &CSVRow{ row: row, }
}

func (t * CSVRow) Size() uint {
	return uint(len(t.row))
}

func (t * CSVRow) GetByIndex(idx uint) (CSVColumn, error) {
	if idx >= t.Size()  {
		return CSVColumn{}, errors.New("Out of range")
	}

	return t.row[idx], nil
}

func (t * CSVRow) GetIndexByValue(value string, ignoreCase bool) int {
	for i := range(t.row) {
		if ignoreCase && (strings.ToLower(value) == t.row[i].ToLower()) {
			return i
		} else if value == t.row[i].value {
			return i
		}
	}

	return -1
}

func (t * CSVRow) Append(c CSVColumn) (* CSVRow) {
	t.row = append(t.row, c)
	return t
}

func (t * CSVRow) EqualSize(u * CSVRow) bool {
	return t.Size() == u.Size()
}

func (t * CSVRow) Apply(callback func(column CSVColumn) CSVColumn) * CSVRow {
	for i := range(t.row) {
		t.row[i] = callback(t.row[i])
	}

	return t;
}

func (t * CSVRow) AsString() []string {
	u := make([]string, 0)
	for i := range(t.row) {
		u = append(u, t.row[i].value)
	}

	return u
}
