package main

import "fmt"

var a int
var b int


//what is that third int for?  that is what makes it run
func plus (a int, b int) int {
          return a + b
      }

func main() {
     fmt.Println("whats are the numbers?")
	 fmt.Scan(&a, &b)       

	 fmt.Println(plus(a, b))

	}

