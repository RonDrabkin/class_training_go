package main

import (
	"fmt"
	"encoding/csv"
	"os"
	)

func main (){

	  f, err:= os.Open("/Users/rondrabkinwikihow/Downloads/tryreading.csv")
  if err != nil {
    fmt.Println("err happened", err)
    } else {
    fmt.Println("I got the file")}
    defer f.Close()


    reader := csv.NewReader(f)

    fmt.Println(reader)
   

	//f.close() - needed?

    //read reads one record from r
    // reader.Read()(f []int, err error) 
}