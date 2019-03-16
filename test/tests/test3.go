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

	fmt.Println("Get column 'sale_date' from first row on index 0")
	fmt.Println(data.GetValueByName(0, "sale_date", true))

	fmt.Println(fmt.Sprintf("Index of 'sale_date' is %d", data.Header().GetIndexByValue("sale_date", true)))

	fmt.Println("Extract index 1,2,4,6,18")

	data1 := data.ExtractIndex(1,2,4,6,18)
	fmt.Println("Get column 'sale_date' from first row on index 0")
	fmt.Println(data1.GetValueByName(0, "sale_date", true))

	fmt.Println("Show first row on index 0")
	fmt.Println(data1.GetRow(0))


}