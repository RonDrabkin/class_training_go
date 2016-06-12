package main

import "fmt"

var age int
func nameage (name string, age int) (string, string, int) {

return name , "is ", age  
}

func main() {

	name := "John"
	age = 27
	fmt.Println(nameage(name, age))


}
	




