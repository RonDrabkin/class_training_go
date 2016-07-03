package main 
//display latitude and longtitude for any given IP address

import ( "fmt"
	   "net/http"
	   "strings"
	   "encoding/json"
	   "ioutil"
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

res, err := http.Get("https://ipfind.co?ip=207.62.246.10&auth=b7e99c21-af76-4b21-bf45-463d019c102d")
if err != nil {
	fmt.Println("oops, error")
}  

body, err := ioutil.ReadAll(res.Body)
TheIP := strings.NewReader((location)resp) //syntax is wrong here

json.NewDecoder(TheIP)
	fmt.Println(TheIP(Latitude), TheIP(Longtitude)) //i didn't see how you println one or two elements of a struct


}