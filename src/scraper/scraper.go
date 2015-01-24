// Abstracts a lot of the scraping stuff
// for ideally any subreddit
package scraper

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

var (
	GFYCAT_SUFFIX = map[string]string{
		"q": "site%3Agfycat.com",
	}

	YOUTUBE_SUFFIX = map[string]string{
		"q": "site%3Ayoutube.com",
	}

	QUERY_SEARCH_PARAMS = map[string]string{
		"restrict_sr": "on",  // restrict subreddit
		"t":           "all", // all time
	}

	SORT_OPTIONS = [...]string{
		"relevance",
		"new",
		"hot",
		"top",
		"comments",
	}
)

// Called during the init, called to fetch all the data
// and store in an abstracted data storage
func scrapeSubreddit(name string) {
	BASE_URL := fmt.Sprintf("reddit.com/r/%s", name)
}

// Given the first page of the page, scrape until
// there is no more next button
func scrapeRoot(url string) {

}
