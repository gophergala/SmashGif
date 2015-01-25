package hello

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"log"
	"net/http"
	"time"

	"appengine"
	"appengine/datastore"
	"appengine/urlfetch"
)

var (
	gifs map[string]Gif
	keys []string
)

func init() {
	m := martini.Classic()
	gifs = make(map[string]Gif)
	// Middle ware stuff
	m.Use(render.Renderer(render.Options{
		Directory: "public",
	}))
	m.Use(martini.Logger())
	m.Use(martini.Static("public"))

	// Request handlers
	m.Get("/", func(r render.Render, req *http.Request) {
		r.HTML(200, "index", nil)
	})

	m.Get("/api", func(r render.Render, req *http.Request) {
		qs := req.URL.Query()
		log.Println(len(keys))
		log.Println("Hitting API endpoint")
		c := appengine.NewContext(req)

		var result []Gif
		q := datastore.NewQuery("Gif").Filter("GifId =", "CheapLeadingBluet")
		_, err := q.GetAll(c, &result)
		check(err)
		//for k, v := range result {
		//	log.Println(k, v, "RESULT IS THIS")
		//}

		gif := queryNext(qs, c)
		log.Println(gif)
		// Return a random gifID for now
		/*gifId := keys[rand.Intn(len(keys))]
		gif := gifs[gifId]*/
		content := gif.Content
		r.JSON(200, map[string]interface{}{
			"id":      gif.GifId,
			"title":   gif.GifTitle,
			"game":    gif.GameTitle,
			"upvotes": content.Upvotes,
			"reddit":  content.Comments,
		})
	})

	m.Get("/scrape", func(res http.ResponseWriter, req *http.Request) {
		c := appengine.NewContext(req)
		client := urlfetch.Client(c)
		var g chan Gif = make(chan Gif)
		// gifs := scrapeSubreddit("smashbros", client)
		go scrapeSubreddit("smashbros", client, g)

		isDone := false
		current := time.Now()

		for !isDone {
			select {
			case gif := <-g:
				storeGif(gif, c)
			case <-time.After(time.Second * 10):
				isDone = true
			}
		}

		after := time.Now()

		keys = extractKeys(gifs)
		log.Println(gifs)
		log.Println("DONE SCRAPING")
		log.Printf("Scraping took %v to run.\n", after.Sub(current))
	})

	http.Handle("/", m)
}
