package main

import "fmt"

func main() {
    
    var fname string
    var lname string

    fmt.Print("What is your name? ")
    fmt.Scan(&fname, &lname)

    fmt.Println(fname, "is a nice first name")

}