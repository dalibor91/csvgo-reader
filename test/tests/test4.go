package main

import (
	"fmt"
	"log"
	"os"
	c "../../csv-reader"
)

func main() {
	file, err1 := os.Open(os.Getenv("test_data"))//os.Open(fmt.Sprintf("%s/sample.csv", dir))
	if err1 != nil {
		log.Fatal(err1)
	}

	k := c.NewCsvReader(file)
	k.SetEndOfLine('\r')
	k.SetEnclosedBy('"')
	k.SetComment('#')
	k.SetDelimiter(',')

	data := k.Read()

	fmt.Println("Show first row on index 0")
	fmt.Println(data.GetRow(0))

	data.Apply(func (row * c.CSVRow) * c.CSVRow {
		return row.Apply(func(column c.CSVColumn) c.CSVColumn {
			column.Set("changed")
			return column
		})
	})

	fmt.Println("Show first row on index 0")
	fmt.Println(data.GetRow(0))


}