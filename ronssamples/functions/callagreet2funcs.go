package main

import "fmt"

var fname string = "first name"
var lname string = "last name"

func main() {

	greet()
	othergreet()
	
	simplegreet()

	fancygreet ("how", "why")
	
}

func greet() {

	fmt.Println("whats your name?")
	fmt.Scan(&fname, &lname)

	fmt.Println("Hello", fname, lname)
}

func othergreet() {
	fmt.Sprint("Hello", fname, lname)
	fmt.Println("Yo ", fname, "baby")
}

func simplegreet () {
	fmt.Println("Whatsup")
}

func fancygreet(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}
