//dont use this is the old badf one

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/openedinc/opened-go"
	"encoding/json"
)

func main() {
    

params := map[string]string{
  "nces_id":"01326677",
  "name":"Pilgrim Lutheran School",
  "address":"8601 Chimney Rock",
  "city": "Houston",
  "state": "TX",
  "zip": "77096",
  "low_grade": "K",
  "high_grade": "8",
}
jsonParams, _ := json.Marshal(params)
client := &http.Client{}
req, _ := http.NewRequest("POST", "https://partner.opened.com/1/schools", &jsonParams)
req.Header.Add("Content-Type", `W/"application/json"`)
token,err := opened.GetToken ("","","","") 
resp, err := client.Do(req)
if err != nil {
    fmt.Println("Errored when sending request to the server")
    return
}
defer resp.Body.Close()
resp_body, _ := ioutil.ReadAll(resp.Body)
fmt.Println(resp.Status)
fmt.Println(string(resp_body))

}
