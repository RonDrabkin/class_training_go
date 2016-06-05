package main

import "fmt"

	var name string

func main() {

	greeting := []string{

		"yo",
		"sock it to me baby",
		"salut",
		//"おはようございます",
	}


	whatsmyname()

    fmt.Println(greeting[2], name)
    greeting = append(greeting, "barf",)
    fmt.Println(greeting)
    fmt.Println(greeting[3])

}

func whatsmyname() {

	
     fmt.Print("What is your name? ")
     fmt.Scan(&name)

}

