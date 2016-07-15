package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 200000; i++ {
		fmt.Printf("the number is %d %b \n", i, i )
	}

}
