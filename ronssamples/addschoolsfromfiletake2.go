package main

import (
	"fmt"
	"github.com/jmcvetta/napping"
	"github.com/openedinc/opened-go"
	"net/http"
  "encoding/csv"
  "bufio"
)




func main() {

	token, err := opened.GetToken("", "", "", "")
	h := &http.Header{}
	h.Add("Authorization", "Bearer " +token)
    
	s := napping.Session{
		Header: h,
	}

  f, _ := os.Open("name of CSV")
  if err != nil {
    fmt.Println("err happened", err)
    } else {
    fmt.Println("I got the file")}

    //all above this tested, compiles and works :)

   defer f.Close()

    school:= map[string]string{
    "nces_id":    "field1",
    "name":       "field2",
    "address":    "field3",
    "city":       "field4",
    "state":      "field5",
    "zip":        "field6",
    "low_grade":  "field7",
    "high_grade": "field8", 

  reader := csv.NewReader(csvfile)

        for {
               record, err := reader.Read()
               //record is a slice of strings, each string is one field
               if err == io.EOF {
                       break
               }
               if err != nil {
                       log.Fatal(err)
               }

               fmt.Println(record)

    //next load each row into the school map

    //rows = school    if rowCount == 0 {
      //for idx, row := range record {
       // columns[column] = idx
     // }
    }
//then marshal into json

  //  jsonParams, _ := json.Marshal(params)
 // client := &http.Client{}

//or do i want to load each row individually and make API requests one by one?
  }
  if err == io.EOF {
      break
  }

    resp, err := s.Post("https://partner.opened.com/1/schools", &school, nil, nil) //return field 1 in 3rd argument?
  if err != nil {
    fmt.Println("Errored when sending request to the server")
    else
    fmt.Println("response status:", resp.Status())
    fmt.Println("school", field1, "was added")
    return

  }



}