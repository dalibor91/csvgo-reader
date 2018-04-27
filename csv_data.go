package csv_reader

type CSVData struct {
	fullData [][]string
	size int
	header []string
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

func (t * CSVData) SetHeader(row []string) * CSVData {
	t.header = row
	return t
}

func (t * CSVData) SetComments(comments []string) * CSVData {
	t.comments = comments
	return t
}

func (t * CSVData) Append(row []string) * CSVData {
	if t.headerLine == (t.size + 1) && !t.headerFound {
		t.headerFound = true
		return t.SetHeader(row)
	}

	t.size = t.size + 1;
	t.fullData = append(t.fullData, row)

	return t
}

func (t * CSVData) Size() int {
	return t.size
}

func (t * CSVData) Data() [][]string {
	return t.fullData
}

func (t * CSVData) Comments() []string {
	return t.comments
}

func (t * CSVData) Header() []string {
	return t.header
}