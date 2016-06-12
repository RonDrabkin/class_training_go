package main 

import "fmt"

func main() {
	
myslice :=  make([]int,1)
fmt.Println(myslice[0])
myslice[0] =7
fmt.Println(myslice[0])
myslice[0] ++
myslice = append(myslice, 800, 400, 300)
fmt.Println(myslice)

//myslice = append(myslice [1:])
//fmt.Println(myslice)

slice2 := myslice[2:3]
fmt.Println(slice2)
}
	
