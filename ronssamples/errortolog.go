package main

import (
	"os"
	"log"
	"fmt"
)

func main() {
	LogFile, err := os.OpenFile("/Users/rondrabkinwikihow/Downloads/log.txt", os.O_WRONLY, 0666)
	//a, err := os.Open("no-file.rtf")
	if err != nil {
			fmt.Println("err happened", err)
			fmt.Println(LogFile)
			//log.SetOutput("/Users/rondrabkinwikihow/Downloads/log.txt")
		    //log.Println("err happened", err)
			//log.Fatalln(err)
       		//panic(err)
	}

		//fmt.Println("whoa, I read the file")
		  log.Println("File Read!")
}
