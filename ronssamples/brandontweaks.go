package main

import "fmt"

func main () {

     var name string = "initial"
     

     fmt.Print("What is your name ")
	 fmt.Scan(&name)

		switch name{

     case "ronald":
     fmt.Println("yes I am the pong king")

 	case "lucas":
 		fmt.Println("Handle These")
 	case "brandon":
 	    fmt.Println("Watch that 2 pointer")	

 	default: 
 	    fmt.Println("that is a nice name")	

        
  
 }}