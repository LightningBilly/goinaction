package search

const defaultName = "default"

type defaultMatcher struct{}

func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}

func init() {
	AddMatcher(defaultName, defaultMatcher{})
}
