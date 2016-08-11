package main
	
import "fmt"


func greeting (firstname string, lastname string) (fullgreeting string) {


	fullgreeting = "hello" + firstname + lastname
	return 
}

func main () {

	var firstname string
	var lastname string
	var fullgreeting string

fmt.Println("what is your first and last name? ")
fmt.Scan(&firstname, &lastname )

greeting (firstname, lastname) 
fmt.Println(fullgreeting)

}