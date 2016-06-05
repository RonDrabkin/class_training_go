package main

import "fmt"

func main() {

	//var x [58]string
	

	y := []string{"a", "b", "luca"}

	for i := 0; i < 50; i++ {

		y = append(y, string(i))
		fmt.Println(y)
	fmt.Println(len(y))
	fmt.Println(cap(y))
	fmt.Println(y[2])
	}

	
	//fmt.Println(x)
	//fmt.Println(len(x))
	//fmt.Println(x[42])

	//for i := 65; i < 122; i++ {

	//	x[i-65] = string(i)


	//}

	//i := "C"

	//x[0] = "Dd"

	//fmt.Println(x)
	//fmt.Println(len(x))

	//fmt.Println(x[0])

}
