package main

import (
	"os"
	"fmt"
)

func main() {
	_, err := os.Open("file name entire path -use control i on right mouse click")
	if err != nil {
		fmt.Println("err happened", err)
		} else {
		fmt.Println("I got the file")}
		//		log.Println("err happened", err)
		//		log.Fatalln(err)
        //		panic(err)
	//data := make([]byte, 100)
//count, err := file.Read(data)
//if err != nil {
//	log.Fatal(err)
}
//fmt.Printf("read %d bytes: %q\n", count, data[:count]) 




// Println formats using the default formats for its operands and writes to standard output.

 