package csv_reader

import "errors"

type CSVData struct {
	fullData []*CSVRow
	size uint
	header *CSVRow
	headerLine int
	headerFound bool
	comments []string
}

func NewCsvData() *CSVData {
	return &CSVData{
		size: 0,
		headerLine: 1,
		headerFound: false,
	}
}

func (t * CSVData) SetHeaderLine(line int) * CSVData {
	t.headerLine = line
	return t
}

func (t * CSVData) setHeader(row []string) * CSVData {
	return t.SetHeader(newCSVRow(row))
}

func (t * CSVData) SetHeader(row * CSVRow) * CSVData {
	t.header = row
	t.headerFound = true
	return t
}

func (t * CSVData) SetComments(comments []string) * CSVData {
	t.comments = comments
	return t
}

func (t * CSVData) Append(row * CSVRow) * CSVData {
	if uint(t.headerLine) == (t.size + 1) && !t.headerFound {
		return t.SetHeader(row)
	}

	t.size = t.size + 1;
	t.fullData = append(t.fullData, row)

	return t
}

func (t * CSVData) append(row []string) * CSVData {
	return t.Append(newCSVRow(row))
}

func (t * CSVData) Size() uint {
	return t.size
}

func (t * CSVData) Data() []*CSVRow {
	return t.fullData
}

func (t * CSVData) Comments() []string {
	return t.comments
}

func (t * CSVData) RemoveRow(index uint) bool {
	if index < t.Size() {
		if (t.Size()-1) == index && t.Size() > 0 {
			t.fullData = t.fullData[:index-1]
			t.size--
		} else {
			t.fullData = append(t.fullData[:index], t.fullData[index+1:]...)
			t.size--
		}
		return true
	}
	return false
}

func (t * CSVData) Header() * CSVRow {
	return t.header
}

func (t * CSVData) GetRow(index uint) (* CSVRow, error) {
	if index >= t.Size() {
		return nil, errors.New("Out of range")
	}

	return t.fullData[index], nil
}

func (t * CSVData) GetHeader() (*CSVRow, error) {
	if !t.headerFound {
		return nil, errors.New("Headers not set")
	}

	return t.header, nil
}

func (t * CSVData) RemoveInvalidRows() uint {
	var removed uint = 0

	for i := range(t.fullData) {
		if !t.fullData[i].EqualSize(t.header) {
			t.RemoveRow(uint(i))
			removed++
		}
	}

	return removed
}

func (t * CSVData) Apply(callback func(row * CSVRow) * CSVRow) * CSVData {
	for i := range(t.fullData) {
		t.fullData[i] = callback(t.fullData[i])
	}

	return t
}

func (t * CSVData) ExtractIndex(idx ... uint) * CSVData {
	data := NewCsvData()

	header := NewCSVRow([]CSVColumn{})

	for i := range idx {
		if row, err := t.header.GetByIndex(idx[i]); err == nil {
			header.Append(row)
		}
	}

	data.SetHeader(header)

	for j := range(t.fullData) {

		tmpRow := NewCSVRow([]CSVColumn{})

		for i := range(idx) {
			if row, err := t.fullData[j].GetByIndex(idx[i]); err == nil {
				tmpRow.Append(row)
			}
		}

		data.Append(tmpRow)
	}

	return data
}

func (t * CSVData) GetValueByName(row uint, name string, ignoreCase bool) (CSVColumn, error) {
	if row, err := t.GetRow(row); err == nil {
		k := t.Header().GetIndexByValue(name, ignoreCase)
		if k != -1 {
			if col, err := row.GetByIndex(uint(k)); err == nil {
				return col, nil
			}
		}
	}
	return CSVColumn{}, errors.New("unable to find")
}

