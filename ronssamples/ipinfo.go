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

resp, err := http.Get("http://api.hostip.info/get_html.php?ip=198.252.210.32&position=true")
if err != nil {
	fmt.Println("oops, error")
} else {

safeIp := url.QueryEscape(resp)

url := fmt.Sprintf("http://api.hostip.info/get_json.php?position=true&ip=%s", safeIp)
fmt.Println(safeIp)
}

fmt.Println("for this IP address, here is the latitude and longtitude", url)

}