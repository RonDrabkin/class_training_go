package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "github.com/openedinc/opened-go"
)

func main() {
  client := &http.Client{}

  token,err := opened.GetToken ("","","","")

  req, _ := http.NewRequest("GET", "https://partner.opened.com/1/publishers.json", nil)

  req.Header.Add("Authorization", "Bearer " +token)

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