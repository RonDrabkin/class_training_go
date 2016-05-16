package main

import "fmt"

  

 func main() {

 	var name string = "initial"

 	fmt.Print("What is your name? ")
    fmt.Scan(&name)
	
	fmt.Println("Hello", name)
}