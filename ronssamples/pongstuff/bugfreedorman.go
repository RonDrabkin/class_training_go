//this thing don't really work, want to try with some kind of if and else
package main

import "fmt"

var name string = "initial"
var score1 int
var score2 int
//var result string

func dormancalcuator(score1 int, score2 int) {

	//if score1 < 10 {

	//	s = fmt.Sprint("nope")
	//	return
	//}
	if score1 > 10 and score1-score2 > 7 {
		fmt.Println("It is a Dorman", name)
		} else {
		fmt.Println("No Dorman this time", name)
		}
 
return}

func main() {

	namechecker()

	fmt.Print("What is the score? ")
	fmt.Scan(&score1)
	fmt.Print(" to ")
	fmt.Scan(&score2)

	dormancalcuator(score1, score2)

	
    

}

func namechecker() {

	fmt.Print("What is your name? ")
	fmt.Scan(&name)

	switch name {

	case "ronald":
		fmt.Println("yes I am the pong king")

	case "lucas":
		fmt.Println("Handle These")
	case "brandon":
		fmt.Println("Watch that 2 pointer")
		//problem here, switch stops when case succeeds, so it doesn't go to the rest of the function now fixed via functions and closure yay

	default:
		fmt.Println("that is a nice name")
	}

}
