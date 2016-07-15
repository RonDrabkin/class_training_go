package main

import (
		"fmt"
	   "math"
	)
var x float64
var Str string

func main () {

	x = 2
	Str = root(x)
	fmt.Println(Str)

}

func root(x float64) string {
	math.Sqrt(x)
	return fmt.Sprint(Str) 
}
