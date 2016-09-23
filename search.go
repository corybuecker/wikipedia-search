package wikipediasearch

import (
	"fmt"
	"net/url"

	"github.com/corybuecker/jsonfetcher"
)

func Search(search string, exactMatch bool) (results []Page, err error) {
	var searchResults = SearchResults{}
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
