package main

import "fmt"

func half(a int) (int, bool) {

	return a / 2, a%2 == 0

}

func main() {
 	a :=2

	var1, var2 := half(a)

	fmt.Println(var1, var2)

}
