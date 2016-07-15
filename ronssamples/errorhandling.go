package main

import "fmt"
	
var x int
func main () {

	x = 2
	str := evalInt(x)
	fmt.Println(str)

}

func evalInt (x int) string {

	if x/2 == 0 {
		return fmt.Sprint("odd")
	} else {

		return fmt.Sprint("even")
	}
}