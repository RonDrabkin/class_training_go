package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Access struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type Schrecord struct {
	School struct {
		ID        json.Number `json:"id,Number"`
		Name      string      `json:"name"`
		NcesID    json.Number `json:"nces_id,Number"`
		LowGrade  json.Number `json:"low_grade,Number"`
		HighGrade json.Number `json:"high_grade,Number"`
		Phone     json.Number `json:"phone,Number"`
		Address   string      `json:"address"`
		City      string      `json:"city"`
		State     string      `json:"state"`
		Zip       json.Number `json:"zip,Number"`
	} `json:"school"`
}

func main() {
	body := []byte("{\n  \"client_id\": \"04c86b71dd20725ff5b54e3654961373561c9909825f7788eeda309b7c14fae5\",\n  \"secret\": \"34c34a4627ecb30003df593ffdac965ecb6f82904e70620f69ef0a9ea2ba8bfe\",\n  \"username\": \"rondrabkin\"\n}")

	req, err := http.NewRequest("POST", "https://partner.opened.com/1/oauth/get_token", bytes.NewBuffer(body))
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal("Errored when sending request to the server")
		return
	}

	defer resp.Body.Close()

	var record Access

	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}


	newbody := []byte("{\n      \"nces_id\": \"umich\",\n      \"name\": \"University of Michigan\",\n      \"address\": \"123 Union Av\",\n      \"city\": \"Ann Arbor\",\n      \"state\": \"MI\",\n      \"zip\": \"95124\",\n      \"low_grade\": \"6\",\n      \"high_grade\": \"12\"\n}")

	newreq, err := http.NewRequest("POST", "https://partner.opened.com/1/schools", bytes.NewBuffer(newbody))
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	token := "Bearer " + record.AccessToken

	newreq.Header.Add("Content-Type", "application/json")
	newreq.Header.Add("Authorization", token)

	newclient := &http.Client{}

	newresp, err := newclient.Do(newreq)

	if err != nil {
		log.Fatal("Errored when sending request to the server")
		return
	}

	defer newresp.Body.Close()

	var newrecord Schrecord

	if err := json.NewDecoder(newresp.Body).Decode(&newrecord); err != nil {
		log.Println(err)
	}

	fmt.Println(newrecord.School)
}
