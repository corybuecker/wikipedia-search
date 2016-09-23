package wikipedia

type Search struct {
	Query struct {
		Pages map[int]struct {
			ID        int    `json:"pageid"`
			Title     string `json:"title"`
			URL       string `json:"canonicalurl"`
			Redirects []struct {
				ID    int    `json:"pageid"`
				Title string `json:"title"`
			} `json:"redirects"`
		} `json:"pages"`
	} `json:"query"`
}
