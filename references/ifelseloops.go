 


package main

import "fmt"

func main() {

	total := 0

	for i := 0; i <= 999; i++ {
		if i%3 == 0 {
			total += i
		} else if i%5 == 0 {
			total += i
		}

	}
	fmt.Println(total)
}

//Here is the solution to https://projecteuler.net/problem=1
//It was actually pretty close to the fizzbuzz one from the prior set of exercises. 
//It is a for loop up to 1000, where you increment the sum by i if i/3 is a whole number, and also increment by i if i/5 is a whole number.  
//the tricky part here was to increment by i only once for multiples of 15, I did that wrong at first ...logic issue not programming :)
