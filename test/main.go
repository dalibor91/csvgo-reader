package main

import (
	c "github.com/dalibor91/csvgo-reader"
	"os"
	"path/filepath"
	"log"
	"fmt"
)


func main() {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	file, err1 := os.Open(fmt.Sprintf("%s/sample.csv", dir))
	if err1 != nil {
		log.Fatal(err1)
	}

	k := c.NewCsvReader(file)
	k.SetEndOfLine('\r')
	k.SetEnclosedBy('"')
	k.SetComment('#')
	k.SetDelimiter(',')
	data := k.Read()

	fmt.Println("Header")
	fmt.Println(data.Header())

	fmt.Println("Data")
	csvdata := data.Data()
	for i := range(csvdata) {
		fmt.Println(csvdata[i])
	}

}