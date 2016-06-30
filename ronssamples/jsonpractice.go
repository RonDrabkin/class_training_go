package main

import (
		"fmt"
		"encoding/json"
		
)

type School struct {

		Name string `json: "school"`
		Address string
		Nces int `json: "nces"`
	}



func main () {
	

	School1 := School{"Design Tech", "Oracle Campus", 12345}
	bs, _ := json.Marshal(School1)
	fmt.Println(bs)
	fmt.Println(string(bs))

	

}

//need a type writer interface {write p[]byte) (n int, err error) }?