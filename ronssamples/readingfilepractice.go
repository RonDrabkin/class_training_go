package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {

	f, err := os.Open("/Users/rondrabkinwikihow/Downloads/tryreading.csv")
	if err != nil {
		fmt.Println("err happened", err)
	} else {
		fmt.Println("I got the file")
	}
	defer f.Close()

	reader := csv.NewReader(f)
	for {

		record, err := reader.Read()

		if err == io.EOF {

			break
		} else if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(record)
	}

}
