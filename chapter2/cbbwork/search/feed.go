package search

import (
	"encoding/json"
	"os"
)

const dataPath = "data/data.json"

type Feed struct {
	Site string `json:"site"`
	Link string `json:"link"`
	Type string `json:"type"`
}

func getFeeds() ([]*Feed, error) {
	file, err := os.Open(dataPath)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	var Feeds []*Feed
	err = json.NewDecoder(file).Decode(&Feeds)

	return Feeds, err
}
