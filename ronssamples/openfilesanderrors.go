package main

import (
	"os"
	"fmt"
	"log"
)

func init () {

	nf, err := os.Create("logger.txt")
	if err != nil {
		fmt.Println(err)
	} //above creates the file logger.txt no problem. now how to write to it.
	log.SetOutput(nf)

}

func main() {
	_, err := os.Open("file name entire path -use control i on right mouse click")
	if err != nil {
		fmt.Println("err happened", err)
		} else {
		fmt.Println("I got the file")
	}
}

		//here we can use log.SetOupt(w io.writer)
		




// Println formats using the default formats for its operands and writes to standard output.

 