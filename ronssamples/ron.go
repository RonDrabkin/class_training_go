package main

//display latitude and longtitude for any given IP address

import (
	"encoding/json"
	"fmt"

	//"log"
	"net/http"
	//"net/url"
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

var IPrecord Location

	ip:= "73.223.170.148"

//safeURL := url.QueryEscape("https://ipfind.co?ip=ip&auth=b7e99c21-af76-4b21-bf45-463d019c102d") 

s := fmt.Sprintf("https://ipfind.co?auth=a3261363-cdc4-47ac-b13e-602445ae7980&ip=%s", ip)

//s := fmt.Sprintf(safeURL , ip)
fmt.Println(s)


//in res we load the json payload
res, err := http.Get(s)
    if err != nil {
       panic(err.Error())
    }

 defer res.Body.Close()   
 fmt.Println(res)

    if err := json.NewDecoder(res.Body).Decode(&record); err != nil {
 	log.Println(err)
 }


       fmt.Println Iprecord.Latitude in the right fomat
 /*but 




    fmt.Println(s.Latitude, s.Longitude)
   

//if err then log fatal  ..dont use else
//defer resp.Body.Close()
	

// Fill the record with the data from the JSON
 //var record IpRecord

 // Use json.Decode for reading streams of JSON data
//then print those values

*/
}
