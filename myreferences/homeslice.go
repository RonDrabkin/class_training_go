package main

import "fmt"

func main() {

	homeslice := []int{1, 3, 5, 7, 9}
	fmt.Println(homeslice)
	fmt.Println(homeslice[3])
	fmt.Println(homeslice[3:])

	madeslice := make([]int, 20, 50)
	fmt.Println(madeslice)

	madeslice[3] = 6
	fmt.Println(madeslice)

	for i := 0; i <= 19; i++ {

		madeslice[i] = i
		fmt.Println(madeslice)
	}

	for i := 20; i <= 50; i++ {

		madeslice = append(madeslice, i)
	}

	fmt.Println(madeslice)

	twodslice := make([][]string, 5)
	fmt.Println(twodslice)
}
