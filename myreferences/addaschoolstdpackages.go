package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/openedinc/opened-go"
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
	
	token,err := opened.GetToken ("","","","")
	if err !=nil {
		fmt.Println("didn't get token")
	}
	token = "Bearer " + token

	req, err := http.NewRequest("POST", "https://partner.opened.com/1/oauth/get_token", bytes.NewBufferString(token))
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


	newbody := []byte("{\n      \"nces_id\": \"abc1234\",\n      \"name\": \"Cooke Elementary School\",\n      \"address\": \"2025 Graves Road,\",\n      \"city\": \"Hockessin\",\n      \"state\": \"DE\",\n      \"zip\": \"19707\",\n      \"low_grade\": \"K\",\n      \"high_grade\": \"5\"\n}")

	newreq, err := http.NewRequest("POST", "https://partner.opened.com/1/schools", bytes.NewBuffer(newbody))
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	

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
