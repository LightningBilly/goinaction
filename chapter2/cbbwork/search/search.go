package search

import (
	"log"
	"sync"
)

var matchers = make(map[string]Matcher)

func Run(searchTerm string) {
	feeds, err := getFeeds()
	if err != nil {
		log.Fatal(err)
	}

	result := make(chan *Result)
	wg := sync.WaitGroup{}
	wg.Add(len(feeds))

	for i, feed := range feeds {
		go func(i int, feed *Feed) {
			feedType := feed.Type
			var mtc Matcher
			if m, ok := matchers[feedType]; ok {
				mtc = m
			} else {
				mtc = matchers[defaultName]
			}
			Match(mtc, feed, searchTerm, result)
			wg.Done()
		}(i, feed)
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	DisplayResult(result)
}

func AddMatcher(feedType string, matcher Matcher) {
	if _, ok := matchers[feedType]; ok {
		log.Fatalln(feedType, "already exists")
	}

	log.Println(feedType, "Add succ")
	matchers[feedType] = matcher
}
