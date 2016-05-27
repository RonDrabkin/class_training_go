package main

import "fmt"
import "strings"

func main() {
    
    var name string

    fmt.Print("What is your name like this Lastname,Firstname? ")
    fmt.Scan(&name)

   Splitname := strings.Split(name, ",")

    for i := range Splitname {
    	fmt.Println(Splitname[i])
    }


}