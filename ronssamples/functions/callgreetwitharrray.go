package main

import "fmt"



func main() {

    var names [3]string 
	fmt.Println("whats your name?")

	fmt.Scanf(&names)

	fmt.Println("Hello", names[1])
}



