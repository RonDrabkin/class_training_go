package main

import "fmt"

func main() {

	var name = 0
	var name2 = 0

	fmt.Print("I will divide for you.  What is the first number? ")
	fmt.Scan(&name)

	fmt.Print("and the second number please?")
	fmt.Scan(&name2)

	fmt.Println(name / name2)
}
