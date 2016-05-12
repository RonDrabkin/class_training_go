package main

import "fmt"

func main () {

     var name string = "initial"
     var score int = "input"
//haven't figured out how to make the integer an input yet
     fmt.Print("What is your name ")
	 fmt.Scan(&name)

		switch name{

     case "ronald":
     fmt.Println("yes I am the pong king")

 	case "lucas":
 		fmt.Println("Handle These")

 	case "brandon":
//    fmt.Print("How many points are you ahead by?")
  //  fmt.Scan(&score)
  //Wanting it to use the variable here to compute an if statement to print out full or half dorman
  //  if score <= 10 {
//        fmt.Println("Full Dorman")
//    } else {
//        fmt.Println("Half Dorman")
  //  }

 	    fmt.Println("Watch that 2 pointer")

 	default:
 	    fmt.Println("that is a nice name")
          }

 }
