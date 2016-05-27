package main

import "fmt"

     var N int 

func main() {

	
	N = Average (10, 21)

	fmt.Println(N)

	}
	

func Average (x int, y int) int {
	return (x + y)/2
}