package main

import (
	"fmt"
	"encoding/csv"
	"os"
	"io"
	"log"
	)

func main (){

	  f, err:= os.Open("/Users/rondrabkinwikihow/Downloads/tryreading.csv")
  if err != nil {
    fmt.Println("err happened", err)
    } else {
    fmt.Println("I got the file")}
    defer f.Close()


    csvReader := csv.NewReader(f)

  for rowCount := 0; ; rowCount++ {
		record, err := csvReader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalln(err)
		}

		columns := make(map[string]int)
		if rowCount == 0 {
			for idx, column := range record {
				columns[column] = idx
			}
		}

		fmt.Println(columns)
		break
	}
   


  
}