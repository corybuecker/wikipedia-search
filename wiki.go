package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/corybuecker/jsonfetcher"
	"github.com/corybuecker/wikipedia-search/wikipedia"
)

func generateURL(game string) string {
	return fmt.Sprintf("https://en.wikipedia.org/w/api.php?action=query&format=json&titles=%s&prop=info|redirects&inprop=url&redirects", url.QueryEscape(game))
}

func getOwnedGames() (*wikipedia.Search, error) {
	var search = &wikipedia.Search{}
	var jsonFetcher = jsonfetcher.Jsonfetcher{}
	if err := jsonFetcher.Fetch(generateURL("Mount & Blade"), &search); err != nil {
		return nil, err
	}

	log.Printf("%+v", search)

	return search, nil
}

func main() {
	getOwnedGames()
}
