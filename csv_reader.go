package csv_reader

import (
	"os"
	"bufio"
)

type CsvReader struct {
	r * os.File
	enclosedBy 	rune
	delimiter 	rune
	endOfLine 	rune
	comment 	rune
	escape 		rune
}

func NewCsvReader(r * os.File) * CsvReader{
	return &CsvReader{
		r: r,
		enclosedBy: '"',
		delimiter: ',',
		endOfLine: '\n',
		comment: '#',
		escape: '\\',
	}
}

func (c * CsvReader) SetEnclosedBy(s rune) * CsvReader {
	c.enclosedBy = s
	return c
}

func (c * CsvReader) SetDelimiter(s rune) * CsvReader {
	c.delimiter = s
	return c
}

func (c * CsvReader) SetEndOfLine(s rune) * CsvReader {
	c.endOfLine = s
	return c
}

func (c * CsvReader) SetComment(s rune) * CsvReader {
	c.comment = s
	return c
}

func (c * CsvReader) Read() * CSVData {

	reader 	:= bufio.NewReader(c.r)
	newLine := true
	lrune	:= '0'

	var cline []string = make([]string, 0);
	var ccol []rune = make([]rune, 0);
	var comm []string = make([]string, 0)

	var data * CSVData = NewCsvData()

	for {
		_rune, _size, _err := reader.ReadRune()
		if _err != nil && _size == 0 {
			break;
		}

		if _rune == c.comment && newLine {
			//handle comment
			var comment []rune = make([]rune, 1)
			comment = append(comment, _rune)
			for {
				_rune1, _size1, _err1 := reader.ReadRune()
				if _err1 != nil && _size1 == 0 && _rune1 != c.endOfLine {
					break;
				}

				comment = append(comment, _rune1)
			}

			comm = append(comm, string(comment))
			newLine = true
			lrune = c.endOfLine
			continue

		} else {
			newLine = false

			if (_rune == c.delimiter || _rune == c.endOfLine || _rune == c.enclosedBy) && lrune == c.escape {
				ccol = append(ccol, _rune)
			} else if _rune == c.delimiter {
				cline = append(cline, string(ccol))
				ccol = make([]rune, 0)
			} else if _rune == c.endOfLine {
				//handle end of line

				cline = append(cline, string(ccol))

				data.append(cline)

				cline = make([]string, 0)
				ccol = make([]rune, 0)
				newLine = true

			} else {
				ccol = append(ccol, _rune)
			}
		}

		lrune = _rune
	}

	data.SetComments(comm)

	return data
}