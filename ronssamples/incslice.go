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

}
	
