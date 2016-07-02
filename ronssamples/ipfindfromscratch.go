package main 
//display latitude and longtitude for any given IP address

import ( "fmt"
	   "net/http"
	   "strings"
	   "encoding/json"
	   )



func main () {


type location struct {
	CountryName string `json:"country"`   
	CountryCode string `json:"country_code"`
	City string `json:"city"`
	IP string `json:"ip_address"`
	Latitude string `json:"latitude"`
	Longtitude string `json:"longtitude"`
}


resp, err := http.Get("GET https://ipfind.co?ip=207.62.246.10&auth=b7e99c21-af76-4b21-bf45-463d019c102d")
if err != nil {
	fmt.Println("oops, error")
} else {

read := strings.NewReader(resp)

json.NewDecoder(read)
	fmt.Println(CountryName)

}
}