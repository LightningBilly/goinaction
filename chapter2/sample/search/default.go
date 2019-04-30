package search

import "fmt"

// defaultMatcher implements the default matcher.
type defaultMatcher struct{}

// init registers the default matcher with the program.
func init() {
	fmt.Println(1)
	var matcher defaultMatcher
	Register("default", matcher)
}

// Search implements the behavior for the default matcher.
func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}
