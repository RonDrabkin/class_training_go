package main
	
import "fmt"


func greeting (firstname string, lastname string) (fullname string) {


	fullname = ("hello" + firstname + lastname)
	return 
}

func main () {

	var firstname string
	var lastname string
	var fullname string

fmt.Println("what is your first and last name? ")
fmt.Scan(&firstname, &lastname )

greeting (firstname, lastname) 
fmt.Println(fullname)

}