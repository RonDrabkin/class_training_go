package main

import (
        "encoding/csv"
        "fmt"
        "io"
        "log"
        "os"
)

func main() {

        csvfile, err := os.Open("/Users/rondrabkinwikihow/Downloads/tryreading.csv")

        if err != nil {
                log.Fatal(err)
        }

        defer csvfile.Close()

        reader := csv.NewReader(csvfile)

        for {
               record, err := reader.Read()
               if err == io.EOF {
                       break
               }
               if err != nil {
                       log.Fatal(err)
               }

               fmt.Println(record)
        }
}

