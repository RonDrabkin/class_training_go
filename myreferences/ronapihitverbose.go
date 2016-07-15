package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// A struct is a collection of fields
// This is our type which matches the JSON object.
type IpRecord struct {
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
	ip := "73.223.170.148"

	url := fmt.Sprintf("https://ipfind.co?auth=a3261363-cdc4-47ac-b13e-602445ae7980&ip=%s", ip)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Get: ", err)
	}

	defer resp.Body.Close()

	var record IpRecord

	// Fill the record with the data from the JSON

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&record)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Latitude = %f and Longitude = %f \n", record.Latitude, record.Longitude)
	fmt.Println("City = ", record.City)
	fmt.Println("Owner = ", record.Owner)
}
