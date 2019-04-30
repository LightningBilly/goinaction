package matchers

import (
    "fmt"
    "log"
    "encoding/xml"
    "errors"
    "net/http"
    "github.com/goinaction/code/chapter2/cbbwork/search"
    "regexp"
)


type (

	item struct {
		XMLName     xml.Name `xml:"item"`
		PubDate     string   `xml:"pubDate"`
		Title       string   `xml:"title"`
		Description string   `xml:"description"`
		Link        string   `xml:"link"`
		GUID        string   `xml:"guid"`
		GeoRssPoint string   `xml:"georss:point"`
	}

    image struct {
        XMLName xml.Name `xml:"image"`
		URL     string   `xml:"url"`
		Title   string   `xml:"title"`
		Link    string   `xml:"link"`
    }

    channel struct {
		XMLName        xml.Name `xml:"channel"`
        Title string `xml:"title"`
        Link string `xml:"link"`
        Description string `xml:"description"`
        PubDate string `xml:"pubDate"`
        LastBuildDate string `xml:"lastBuildDate"`
		Language       string   `xml:"language"`
        Image image `xml:"image"`
        Item    []item  `xml:"item"`
    }

    rssDocument struct {
		XMLName xml.Name `xml:"rss"`
        Channel channel `xml:"channel"`
    }
)

type rssMatcher struct {}

func (m rssMatcher) Search(feed *search.Feed, searchTerm string) ([]*search.Result, error) {
    log.Printf("search Url[%s]", feed.Link)
    resp, err := getHttp(feed.Link)
    if err != nil {
        return nil, err
    }

    var results []*search.Result
    for _, channelItem := range resp.Channel.Item {
    	//log.Println(channelItem.Title)
        matched, err := regexp.MatchString(searchTerm, channelItem.Title)
        if err != nil {
            return nil, err
        }
        if matched {
            results = append(results, &search.Result{
                Field: "Title",
                Content: channelItem.Title,
            })
        }
    }

    return results, nil
}

func getHttp(url string) (*rssDocument, error) {
    if url == "" {
        return nil, errors.New("url is empty")
    }

    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    if resp.StatusCode != 200 {
        return nil, fmt.Errorf("HTTP Response Error %d\n", resp.StatusCode)
    }
    var doc rssDocument
    err = xml.NewDecoder(resp.Body).Decode(&doc)
    return &doc, err
}

func init() {
    var matcher rssMatcher
    search.AddMatcher("rss", matcher);
}
