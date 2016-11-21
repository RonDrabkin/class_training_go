//downloads title tag and meta

package main

import (
    // import standard libraries
    "fmt"
    "log"
    "net/http"

    // import third party libraries
    "github.com/PuerkitoBio/goquery"
)

func metaScrape() {
    doc, err := goquery.NewDocument("http://elementymcacamp.com/")
    if err != nil {
        log.Fatal(err)
    }

    var metaDescription string
    var pageTitle string
    

    // use CSS selector found with the browser inspector
    // for each, use index and item
    pageTitle = doc.Find("title").Contents().Text()
   
    
       
    

    doc.Find("meta").Each(func(index int, item *goquery.Selection) {
        if( item.AttrOr("name","") == "description") {
            metaDescription = item.AttrOr("content", "")
        }
    })
    fmt.Printf("Page Title: '%s'\n", pageTitle)
    fmt.Printf("Meta Description: '%s'\n", metaDescription)
    
  
}

func main() {
    metaScrape()
    resp, err :=http.Get("http://elementymcacamp.com/")
    if err !=nil {
        fmt.Println(err)
    }
    fmt.Println("StatusCode:", resp.StatusCode)
}