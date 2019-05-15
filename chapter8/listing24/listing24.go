// This sample program demonstrates how to decode a JSON response
// using the json package and NewDecoder function.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type (
	// gResult maps to the result document received from the search.
	gResult struct {
		GsearchResultClass string `json:"GsearchResultClass"`
		UnescapedURL       string `json:"unescapedUrl"`
		URL                string `json:"url"`
		VisibleURL         string `json:"visibleUrl"`
		CacheURL           string `json:"cacheUrl"`
		Title              string `json:"title"`
		TitleNoFormatting  string `json:"titleNoFormatting"`
		Content            string `json:"content"`
	}

	// gResponse contains the top level document.
	gResponse struct {
		ResponseData struct {
			Results []gResult `json:"results"`
		} `json:"responseData"`
	}
)

var JSON = `{
	"responseData": {
		"results": [{
				"GsearchResultClass": "GwebSearch",
				"unescapedUrl": "https://www.reddit.com/r/golang",
				"url": "https://www.reddit.com/r/golang",
				"visibleUrl": "www.reddit.com",
				"cacheUrl": "http://www.google.com/search?q=cache:W...",
				"title": "r/\u003cb\u003eGolang\u003c/b\u003e - Reddit",
				"titleNoFormatting": "r/Golang - Reddit",
				"content": "First Open Source \\u003cb\\u003eGolang\\u...",
				"num": 10000
			},
			{
				"GsearchResultClass": "GwebSearch",
				"unescapedUrl": "http://tour.golang.org/",
				"url": "http://tour.golang.org/",
				"visibleUrl": "tour.golang.org",
				"cacheUrl": "http://www.google.com/search?q=cache:O...",
				"title": "A Tour of Go",
				"titleNoFormatting": "A Tour of Go",
				"content": "Welcome to a tour of the Go programming ..."
			}
		]
	}
}`

func main() {
	uri := "http://ajax.googleapis.com/ajax/services/search/web?v=1.0&rsz=8&q=golang"
	uri = "http://localhost:1323/search"

	// Issue the search against Google.
	resp, err := http.Get(uri)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}
	defer resp.Body.Close()
	var b []byte
	fmt.Println(resp.Body.Read(b))

	log.Println("body :", string(b))

	// Decode the JSON response into our struct type.
	var gr gResponse
	err = json.NewDecoder(resp.Body).Decode(&gr)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println(gr)

	// Marshal the struct type into a pretty print
	// version of the JSON document.
	pretty, err := json.MarshalIndent(gr, "", "    ")
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println(string(pretty))

	err = json.Unmarshal([]byte(JSON), &h)
}
