package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/openedinc/opened-go"
	"io"
	"log"
	"encoding/csv"
	"bufio"
	"net/http"
	"os"
)

type Access struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

//remove the below, or at least, remove the wrapper

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

	token, err := opened.GetToken("", "", "", "")
	if err != nil {
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

	f, err := os.Open("/Users/rondrabkinwikihow/Downloads/addschooltest.csv")

	r := csv.NewReader(bufio.NewReader(f))
	for {
		record, err := r.Read()

		if err == io.EOF {

			break
		}

		fmt.Println(record)
		fmt.Println(len(record))
		//for value := range record {
		//	fmt.Printf("  %v\n", record[value])
		fmt.Println(record[0])
		nces_id:=record[0]
		fmt.Println(record[1])
		schoolname:=record[1]
		address:=record[2]
		city:=record[3]
		state:=record[4]
		zip:=record[5]
		lowgrade:=record[6]
		highgrade:=record[7]

		}
	}


newbody := []byte("{\n      \"nces_id\": \"2400510\",\n      \"name\": \"Prince George's County Public Schools\",\n      \"address\": \"14201 School Lane\",\n      \"city\": \"Upper Marlboro\",\n      \"state\": \"MD\",\n      \"zip\": \"20772\",\n      \"low_grade\": \"K\",\n      \"high_grade\": \"12\"\n}")
/*
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
	fmt.Println("Successfully added ", newrecord.School)
}
}
}
}
*/