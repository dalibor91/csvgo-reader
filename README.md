# csvgo-reader
CSV Parser for folang 


Csv reader that enables you to specify terminating char, delimiter, comments , enclosure, escape string 

Please check test for example how to use it 

https://github.com/dalibor91/csvgo-reader/blob/master/test/main.go

```
package main

import (
        c "github.com/dalibor91/csvgo-reader"
        "os"
        "log"
        "fmt"
)


func main() {
        file, err1 := os.Open("file.csv")
        if err1 != nil {
                log.Fatal(err1)
        }

        k := c.NewCsvReader(file)
        k.SetEndOfLine('\n')
        k.SetEnclosedBy('"')
        k.SetComment('$')
        k.SetDelimiter('|')
        data := k.Read()

        fmt.Println("Header")
        fmt.Println(data.Header())

        fmt.Println("Data")
        csvdata := data.Data()
        for i := range(csvdata) {
                fmt.Println(csvdata[i])
        }

}
```
