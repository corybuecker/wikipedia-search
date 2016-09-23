package wikipedia

type Search struct {
	Page  Page
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

func (search *Search) All() []Page {
	var pages = make([]Page, 0, len(search.Query.Pages))

	for _, page := range search.Query.Pages {
		pages = append(pages, page)
	}

	return pages
}

func (search *Search) Filtered(searchValue string) []Page {
	var pages = make([]Page, 0, len(search.Query.Pages))

	for _, page := range search.Query.Pages {
		if page.Title == searchValue {
			pages = append(pages, page)
		}
	}

	return pages
}
