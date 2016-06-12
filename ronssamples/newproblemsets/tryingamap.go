package main

import "fmt"

func main() {

var greetingmap = make(map[string]string)

	greetingmap["Adin"] = "how is the game"
	greetingmap["Hiroshi"] = "オス"
	greetingmap["Davina"] = "hi cutie"

	fmt.Println(greetingmap)
	fmt.Println(len(greetingmap))
	fmt.Println(greetingmap["Davina"])


}
	




