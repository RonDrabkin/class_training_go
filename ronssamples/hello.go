package main

import  (
		"fmt"
		"bufio"
		"os"
		)
func main() {


	fmt.Print("Please enter your name: ")
	scanner := bufio.NewScanner(os.Stdin)
            for scanner.Scan() {
    fmt.Println("hello", scanner.Text())
    
    break
               

	
}

}

