package main 
//display latitude and longtitude for any given IP address

import ( "fmt"
	   "net/url"
	   "net/http"
	   )



func main () {

//fmt.Println("Please enter the IP addy"  ....do this part later
// create a new request, a new client and use Do function	
//https://mholt.github.io/json-to-go/	

//create a user defined type to store the data from the API pull

type location struct {
	CountryName string `json:"country_name"`   //that is called a tag.  json is what used to be xml sorta
	CountryCode string `json:"country_code"`
	City string `json:"city"`
	IP string `json:"ip"`
	Latitude string `json:"latitude"`
	Longtitude string `json:"longtitude"`
}

//query the API to get the data on the location
resp, err := http.Get("https://ipfind.co?ip=73.223.170.148&auth=b7e99c21-af76-4b21-bf45-463d019c102d")
if err != nil {
	fmt.Println("oops, error")
} else {

safeIp := url.QueryEscape(resp)  
//escape the query to use encoding
//error cannot use resp (type *http.Response) as type string in argument to url.QueryEscape
//i tried resp string, the example is s string
//my class says you need to decode . but in this case must query.escape first to decode it?

url := fmt.Sprintf("http://api.hostip.info/get_json.php?position=true&ip=%s", safeIp)
//Sprint returns as a string instead of printing etc....why are we getting json/string
//i'm confused why return the html and then return the json.  why not just do one the json?  or is it in the API docs?

fmt.Println(safeIp)
}

//wamt to put the returned json into the custom type so we can print it, that is where I'm confused, how to get from 
//returned json into the fields of the custom type that must be in the next lesson
mylocation := location(CountryName, CountryCode, City, IP, Latitude, Longtitude)

//i think the below is how you print an element of a struct
fmt.Print(mylocation.CountryName, mylocation.CountryCode, mylocation.City, mylocation.IP, "is at ")
fmt.Println(mylocation.Latitude, mylocation.Longtitude)

}