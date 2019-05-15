package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type (
	Result struct {
		GsearchResultClass string `json:"GsearchResultClass"`
		UnescapedUrl       string `json:"unescapedUrl"`
		Num                int    `json:"num"`
	}
	hResponse struct {
		ResponseData struct {
			Results []Result `json:"results"`
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
	uri := "http://localhost:1323/search"
	resp, err := http.Get(uri)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}

	defer resp.Body.Close()

	var h hResponse
	err = json.NewDecoder(resp.Body).Decode(&h)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println(h)

	pretty, err := json.MarshalIndent(h, "---", "    ")
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}

	fmt.Println(string(pretty))

	err = json.Unmarshal([]byte(JSON), &h)
	if err != nil {
		fmt.Println("ERROR:", err)
		return

	}

	pretty, err = json.MarshalIndent(h, "", "    ")
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}

	fmt.Println(string(pretty))
	io.Writer()
}
