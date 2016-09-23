package main

import (
	"fmt"
	"net/url"

	"github.com/corybuecker/jsonfetcher"
	"github.com/corybuecker/wikipedia-search/wikipedia"
)

func Search(search string, exactMatch bool) (results []wikipedia.Page, err error) {
	var searchResults = wikipedia.Search{}
	var jsonFetcher = jsonfetcher.Jsonfetcher{}

	if err := jsonFetcher.Fetch(generateURL(search), &searchResults); err != nil {
		return nil, err
	}

	if exactMatch {
		return searchResults.Filtered(search), nil
	}

	return searchResults.All(), nil
}

func generateURL(search string) string {
	return fmt.Sprintf("https://en.wikipedia.org/w/api.php?action=query&format=json&titles=%s&prop=info|redirects&inprop=url&redirects", url.QueryEscape(search))
}
