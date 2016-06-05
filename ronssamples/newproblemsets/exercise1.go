package main
	
import "fmt"

   var x int

func main () {

 	fmt.Print("what is the number? ")
 	fmt.Scan(&x)

 	fmt.Println(x/2)
 	if (x%2) != 0 {
	fmt.Println("its odd") 
	} else {
    fmt.Println("its even")}




	
//func one x(int) 	
//	fmt.Println(x/2)

//func two x(int)
//	if x/2 ret != 0 {
//	fmt.Println("its odd") 
//	} else {fmt.Println("its true")}




}