
package main

import (
	"log"
	"fmt"
	"math"
)

	var number float64
	var number1 float64

func main() {


	fmt.Print("What is the number? ")
	fmt.Scan(&number)

	num, err := Sqrt(number)
	if err != nil {
		log.Fatalln(err)
	} else if num >= 0 {
	
		number1 := math.Sqrt(number)
		fmt.Println("the square root is", number1)
	}
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("oh boy you should know this: square root of negative number: %v", f)
	}
	// implementation
	
	return number, nil
}