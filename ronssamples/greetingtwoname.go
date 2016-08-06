package main
	
import "fmt"


func greeting (firstname string, lastname string) string {

	return fmt.Sprintf("hello", firstname, lastname)
}

func main () {

	var firstname string
	var lastname string

fmt.Println("what is your first and last name? ")
fmt.Scan(&firstname, &lastname )

greeting (firstname, lastname) 
}