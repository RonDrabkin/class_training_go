package main

import "fmt"

func main() {

	var score1 int
	var score2 int

	fmt.Println("What is the score?")
	fmt.Scan(&score1)
	fmt.Println(" to ")
	fmt.Scan(&score2)

	if score1-score2 > 10 {

		fmt.Println("I can't calculate a Dorman, but you are really getting spanked")

	} else {

		fmt.Println("Great, looks like a close one!")
	}

}
