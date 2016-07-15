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

/*
func evalInt (x int) string {

	if x/2 == 0 {
		return fmt.Sprint("odd")
	} else {

		return fmt.Sprint("even")
	}
}
*/
func root(x float64) string {
	math.Sqrt(x)
	return fmt.Sprint(Str) 
}
