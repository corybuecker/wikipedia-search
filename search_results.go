package wikipediasearch

type SearchResults struct {
	Query struct {
		Pages map[string]Page `json:"pages"`
	} `json:"query"`
}

type Page struct {
	ID        int    `json:"pageid"`
	Title     string `json:"title"`
	URL       string `json:"canonicalurl"`
	Redirects []struct {
		ID    int    `json:"pageid"`
		Title string `json:"title"`
	} `json:"redirects"`
}

func (search *SearchResults) All() []Page {
	var pages = make([]Page, 0, len(search.Query.Pages))

	for key, page := range search.Query.Pages {
		if key != "-1" {
			pages = append(pages, page)
		}
	}

	return pages
}

func (search *SearchResults) Filtered(searchValue string) []Page {
	var pages = make([]Page, 0, len(search.Query.Pages))

	for key, page := range search.Query.Pages {
		if key != "-1" && pageMatch(page, searchValue) {
			pages = append(pages, page)
		}
	}

	return pages
}

func pageMatch(page Page, search string) bool {
	if page.Title == search {
		return true
	}

	for _, redirect := range page.Redirects {
		if redirect.Title == search {
			return true
		}
	}

	return false
}
