package main
	
import "fmt"



  func half(x int) (int){
   
    return x/2
 }
 

func main () {
	var x int

 	fmt.Print("what is the number? ")
 	fmt.Scan(&x)

    x= half(x)
    fmt.Println(x)

    pow := make([]int, 10)
    for i:= range pow {
    	pow[i] = i
    fmt.Println(pow)	
    }


}




	
//func one x(int) 	
//	fmt.Println(x/2)

//func two x(int)
//	if x/2 ret != 0 {
//	fmt.Println("its odd") 
//	} else {fmt.Println("its true")}




