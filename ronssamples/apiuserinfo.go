package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	//"github.com/openedinc/opened-go"
)

func main() {
	client := &http.Client{}

	req, _ := http.NewRequest("GET", "https://partner.opened.com/1/resources.json/?descriptive=World War II", nil)

	//token,err := opened.GetToken ("","","","")

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
}
