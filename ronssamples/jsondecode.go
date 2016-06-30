package main

import (
		"fmt"
		"encoding/json"
		"strings"
)

type School struct {

		Name string `json: "school"`
		Address string
		Nces int `json: "nces"`
	}



func main () {
	var p1 School 


	read := strings.NewReader(`{"Name":"Design Tech","Address":"Oracle Campus","Nces":12345}`)
	
	json.NewDecoder(read), Decode(&p1)
	fmt.Println(p1,Name)

}

//need a type writer interface {write p[]byte) (n int, err error) }?