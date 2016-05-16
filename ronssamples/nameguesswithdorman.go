package main

import "fmt"

func main() {

	var name string = "initial"
	var score1 int
	var score2 int

	fmt.Print("What is your name ")
	fmt.Scan(&name)

	switch name {

	case "ronald":
		fmt.Println("yes I am the pong king")

	case "lucas":
		fmt.Println("Handle These")
	case "brandon":
		fmt.Println("Watch that 2 pointer")

	default:
		fmt.Println("that is a nice name")

		fmt.Println("What is the score?")
		fmt.Scan(&score1)
		fmt.Println(" to ")
		fmt.Scan(&score2)

		if score1-score2 > 10 {

			fmt.Println(name, "I can't calculate a Dorman, but you are really getting spanked")

		} else {

			fmt.Println("Great, looks like a close one,", name)
		}

	}
}
