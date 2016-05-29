package main

import "fmt"
import "strings"

var name string = "initial"
var score1 int
var score2 int
var result string

func dormancalcuator(score1 int, score2 int) (result string) {

	if score1-score2 > 7 && score1 > 10 {
		fmt.Print("It is a Dorman, ")
		fmt.Println(strings.Title(name))
	} else {
		fmt.Print("No Dorman this time, ")
		fmt.Println(strings.Title(name))
	}

	return
}

func main() {

	namechecker()

	fmt.Print("How many points to you have? ")
	fmt.Scan(&score1)
	fmt.Print("How many points does your opponent have? ")
	fmt.Scan(&score2)

	dormancalcuator(score1, score2)

}

func namechecker() {

	fmt.Print("What is your name? ")
	fmt.Scan(&name)

	name=(strings.ToLower(name))
    //you'd think there is a better way to do this case insenstive thing

	switch name {

	case "ronald":
		fmt.Println("yes I am the pong king")

	case "lucas":
		fmt.Println("Handle These")
	case "brandon":
		fmt.Println("Watch that 2 pointer")

	default:
		fmt.Println("that is a nice name")
	}

}
