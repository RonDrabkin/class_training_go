package main

//display latitude and longtitude for any given IP address

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Location struct {
	IPAddress     string  `json:"ip_address"`
	Country       string  `json:"country"`
	CountryCode   string  `json:"country_code"`
	Continent     string  `json:"continent"`
	ContinentCode string  `json:"continent_code"`
	City          string  `json:"city"`
	County        string  `json:"county"`
	Region        string  `json:"region"`
	RegionCode    string  `json:"region_code"`
	Timezone      string  `json:"timezone"`
	Owner         string  `json:"owner"`
	Longitude     float64 `json:"longitude"`
	Latitude      float64 `json:"latitude"`
}

func main() {

var res Location

	ip:= "73.223.170.148"

safeURL := url.QueryEscape("https://ipfind.co?ip=ip&auth=b7e99c21-af76-4b21-bf45-463d019c102d") 

s := fmt.Sprintf(safeURL, "string")

//sprintF with IP
//http:Get only 

res, err := http.Get("safeURL")
    if err != nil {
        panic(err.Error())
    }

    body, err := ioutil.ReadAll(res.Body)

    if err != nil {
        panic(err.Error())
    }
       
 defer res.Body.Close()


    json.Unmarshal(body, &res)

    fmt.Println(res{Latitude, Longitude})

//if err then log fatal  ..dont use else
//defer resp.Body.Close()
	

// Fill the record with the data from the JSON
 //var record IpRecord

 // Use json.Decode for reading streams of JSON data
//then print those values


}
