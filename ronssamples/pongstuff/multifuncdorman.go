//this thing don't really work, want to try with some kind of if and else
package main

import "fmt"

var name string = "initial"
var score1 int
var score2 int
var s string

func dormancalcuator(score1 int, score2 int) (s string) {

	//if score1 < 10 {

	//	s = fmt.Sprint("nope")
	//	return
	//}
	if score1-score2 > 7 {
		s = "Its a Dorman!"
		} else {
		s = "Sorry No Dorman"
		}
 
return}

func main() {

	namechecker()

	fmt.Print("What is the score? ")
	fmt.Scan(&score1)
	fmt.Println(" to ")
	fmt.Scan(&score2)

	dormancalcuator(score1, score2)

	
       fmt.Println(name, s)

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
