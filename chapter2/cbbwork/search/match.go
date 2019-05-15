package search

import (
	"log"
)

type Result struct {
	Field   string
	Content string
}

type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

func Match(matcher Matcher, feed *Feed, searchTerm string, result chan<- *Result) {
	log.Println(*feed, searchTerm)
	res, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println("search err", err)
		return
	}

	for _, r := range res {
		result <- r
	}
}

func DisplayResult(result <-chan *Result) {
	for res := range result {
		log.Printf("%s\n\n%s\n\n", res.Field, res.Content)
	}
}
