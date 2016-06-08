package main

import "fmt"

func main() {
	
var nextone int
var eventotal int

	 fibo := []int{1,2} 

for i := 2 ; nextone<=3500000 ; i++  {

	nextone = fibo[i-1] + fibo[i-2] 

	fibo = append(fibo, nextone)
    }

		fmt.Println(fibo)
		fmt.Println(len(fibo))
	

 
  for i := 0 ; i <= 31; i++ {

  		if fibo[i] % 2  == 0 {
  		eventotal = eventotal + fibo[i]
    }

  }

  fmt.Println(eventotal)
	
}


