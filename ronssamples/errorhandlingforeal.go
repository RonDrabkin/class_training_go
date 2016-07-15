package main

import (
		"fmt"
	   "math"
	)
var x float64
var Answer float64

func main () {

	x = 2
	Answer = root(x)
	fmt.Println(Answer)

}

func root(float64) float64 {
	
	return math.Sqrt(x)

}
