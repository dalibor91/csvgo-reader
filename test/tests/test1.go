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

	header := data.Header()
	comments := data.Comments()

	fmt.Println("Following headers found:")
	var i uint = 0;
	for ; i < header.Size(); i++ {
		col, _ := header.GetByIndex(i)
		fmt.Println(col.Get())
	}

	fmt.Println("\n\n\n\nComments: ")
	for j := range(comments) {
		fmt.Println(comments[j])
	}
}