package main

import (
	//c "github.com/dalibor91/csvgo-reader"
	c "../../csv-reader"
	"fmt"
	"log"
	"os"
)


func main() {

	/*dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}*/



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

	/*for {
		fmt.Print("> ")
	}*/

	fmt.Println("Header")
	fmt.Println(data.Header())

	//fmt.Println("Data")
	//csvdata := data.Data()
	//for i := range(csvdata) {
	//	fmt.Println(csvdata[i])
	//}

	fmt.Println(fmt.Sprintf("Found %d rows", data.Size()))

	irows := data.RemoveInvalidRows()

	fmt.Println(fmt.Sprintf("Removed %d invalid rows", irows))

	fmt.Println("Show first row on index 0")
	fmt.Println(data.GetRow(0))
	fmt.Println("After apply")
	row, _ := data.GetRow(0)
	row.Apply(func(column c.CSVColumn) c.CSVColumn {
		column.Set("keks")
		return column
	});
	fmt.Println(data.GetRow(0))

	fmt.Println("Get column 'sale_date' from first row on index 0")
	fmt.Println(data.GetValueByName(0, "sale_date", true))

	fmt.Println(fmt.Sprintf("Index of 'sale_date' is %d", data.Header().GetIndexByValue("sale_date", true)))


	fmt.Println("Extract index 1,2,4,6,18")

	data.Apply(func (row * c.CSVRow) * c.CSVRow {
		return row.Apply(func(column c.CSVColumn) c.CSVColumn {
			column.Set("changed")
			return column
		})
	})

	data1 := data.ExtractIndex(1,2,4,6,18)

	fmt.Println(fmt.Sprintf("Found %d rows", data1.Size()))

	fmt.Println("Header")
	fmt.Println(data1.Header())






}