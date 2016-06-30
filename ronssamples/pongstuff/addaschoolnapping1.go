package main

import (
	"fmt"
	"github.com/jmcvetta/napping"
	"github.com/openedinc/opened-go"
	"net/http"
)

func main() {

	token, err := opened.GetToken("", "", "", "")
	h := &http.Header{}
	h.Add("Authorization", "Bearer " +token)
    
	s := napping.Session{
		Header: h,
	}

	payload := map[string]string{
		 "nces_id":"noneyet",
  "name":"iLead Pacoima",
  "address":"11261 Glenoaks Blvd",
  "city": "Pacoima",
  "state": "CA",
  "zip": "91331",
  "low_grade": "K",
  "high_grade": "12",
	}

	resp, err := s.Post("https://partner.opened.com/1/schools", &payload, nil, nil)
	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return

		//jsonParams, _ := json.Marshal(params)
		////req, _ := http.NewRequest("POST", "https://partner.opened.com/1/schools", &jsonParams)
		//req.Header.Add("Content-Type", `W/"application/json"`)

	}
	//defer resp.Body.Close()
	//resp_body , _:= ioutil.ReadAll(napping.body)
	fmt.Println("response status:", resp.Status())
	//fmt.Println(string(resp_body))

}
