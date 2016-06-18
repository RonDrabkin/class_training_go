package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	client := &http.Client{}

	body := []byte({\n      \"nces_id\":\"BB981479\",\n      \"name\":\"Woodlands Christian Academy\",\n      \"address\":\"5900 Academy Way\",\n      \"city\": \"The Woodlands\",\n      \"state\": \"TX\",\n      \"zip\": \"77384\",\n      \"low_grade\": \"K\",\n      \"high_grade\": \"12\"\n    }\n})

	req, _ := http.NewRequest("POST", "https://partner.opened.com/1/schools", bytes.NewBuffer(body))

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer [TOKEN]")

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
  
    

