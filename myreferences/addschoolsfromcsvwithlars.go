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

type School struct {
	Name      string
	NcesID    string
	LowGrade  string
	HighGrade string
	Phone     string
	Address   string      
	City      string      
	State     string      
	Zip       string
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
	for i := 0; i < 2; i ++ {
		fmt.Printf("i: %d\n", i)
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		// fmt.Println(record)
		// fmt.Println(len(record))
		// //for value := range record {
		//	fmt.Printf("  %v\n", record[value])
		//fmt.Println(record[0])
		// NCES:=record[0]
		// fmt.Println(record[1])
		// schoolname:=record[1]

		school := School{
			Name: record[0],
			NcesID: record[1],
			Address: record[2],
			City:    record[3],
			State:   record[4],
			Zip:	 record[5],
			LowGrade: record[6],
			HighGrade: record[7],
		}
		//fmt.Printf("Record: %+v\n", record)

		fmt.Printf("School: %+v\n", school)
	}
}


//newbody := []byte("{\n      \"nces_id\": \"2400510\",\n      \"name\": \"Prince George's County Public Schools\",\n      \"address\": \"14201 School Lane\",\n      \"city\": \"Upper Marlboro\",\n      \"state\": \"MD\",\n      \"zip\": \"20772\",\n      \"low_grade\": \"K\",\n      \"high_grade\": \"12\"\n}")
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
