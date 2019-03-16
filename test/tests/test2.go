package main

import (
	"fmt"
	"log"
	"os"
	c "../../../csvgo-reader"
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

	fmt.Println(fmt.Sprintf("Found %d rows", data.Size()))

	irows := data.RemoveInvalidRows()

	fmt.Println(fmt.Sprintf("Removed %d invalid rows", irows))

	fmt.Println("Show first row on index 0")
	fmt.Println(data.GetRow(0))
	fmt.Println("After apply")
	row, _ := data.GetRow(0)
	row.Apply(func(column c.CSVColumn) c.CSVColumn {
		column.Set("apply")
		return column
	});
	fmt.Println(data.GetRow(0))

}