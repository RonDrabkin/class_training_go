package main

import "fmt"

func main () {

     var fname string
     var lname string

     fmt.Println("whats your name?")
     fmt.Scan(&fname, &lname)
     fmt.Println("Hello ", fname, lname)
 }