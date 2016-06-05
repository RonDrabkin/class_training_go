package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/csv"
	"log"
)

func main() {
	client := &http.Client{}

	req, _ := http.NewRequest("GET", "http://polls.apiblueprint.org/questions", nil)

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))

	 w := csv.NewWriter(os.Stdout)
	  for _, record := range records {
        if err := w.Write(record); err != nil {
            log.Fatalln("error writing record to csv:", err)
        }

        w.Flush()
//take output in the terminal and paste into a json viewer
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
}