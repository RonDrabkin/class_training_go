package main

import (
		"fmt"
		"bufio"
		"os"
		"in"
		)

func main() {
    fmt.Print("whats your name: ")
   
    name := bufio.NewReader(os.Stdin)
    line, err := in.ReadString('\n')


   fmt.Println(name)
}