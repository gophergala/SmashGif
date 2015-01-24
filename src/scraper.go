// Abstracts a lot of the scraping stuff
// for ideally any subreddit
package hello

import (
	"appengine"
	"appengine/urlfetch"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
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

func prepareUrl(base string) string {
	var queryParams []string
	for k, v := range QUERY_SEARCH_PARAMS {
		queryParams = append(queryParams, fmt.Sprintf("%s=%s", k, v))
	}
	return fmt.Sprintf("%s?%s", base, strings.Join(queryParams, ""))
}

// Called during the init, called to fetch all the data
// and store in an abstracted data storage
func scrapeSubreddit(name string) map[string]Gif {
	BASE_URL := fmt.Sprintf("https://reddit.com/r/%s", name)
	req, err := http.NewRequest("GET", BASE_URL, nil)
	check(err)
	c := appengine.NewContext(req)
	client := urlfetch.Client(c)

	return scrapeRoot(prepareUrl(BASE_URL), client)
}

// Given the first page of the page, scrape until
// there is no more next button
func scrapeRoot(url string, c *http.Client) map[string]Gif {
	gifs = make(map[string]Gif)
	for nextUrl := url; nextUrl != ""; {
		pageGifs, temp := scrapePage(nextUrl, c)
		nextUrl = temp
		extendMap(gifs, pageGifs)
	}

	return gifs
}

func scrapePage(url string, client *http.Client) (map[string]Gif, string) {
	resp, err := client.Get(url)
	check(err)

	doc, err := goquery.NewDocumentFromResponse(resp)
	check(err)

	var gifs = make(map[string]Gif)
	doc.Find(".linkflair").Each(func(i int, s *goquery.Selection) {
		link, exists := s.Find("a.thumbnail").Attr("href")
		if !exists {
			return
		}

		// Gets all of the data
		votes, err := strconv.ParseUint(s.Find("div.score.unvoted").Text(), 10, 32)
		check(err)

		titles := s.Find("p.title").Children().First()
		comments, exists := s.Find(".comments").Attr("href")
		if !exists {
			return
		}

		// We're targetting smash bros for now
		gameTitle := titles.Text()
		gifTitle := titles.Next().Text()

		re := regexp.MustCompile(`^https?:\/\/[a-z\:0-9.]+\/`)
		link = re.ReplaceAllString(link, "")
		gifId := strings.Split(link, "?")[0]

		gifs[gifId] = Gif{
			Content{
				comments:  comments,
				upvotes:   votes,
				subreddit: "smashbros",
			},
			gameTitle,
			gifTitle,
			gifId,
		}
		log.Println(gifs[gifId])
	})

	nextUrl := ""
	doc.Find("span.nextprev").Children().Each(func(i int, s *goquery.Selection) {
		if strings.Contains(s.Text(), "next") {
			nextUrl, _ = s.Attr("href")
		}
	})

	return gifs, nextUrl
}