package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 20000000; i++ {
		fmt.Printf("the number is %d %b \n", i, i )
	}

}
