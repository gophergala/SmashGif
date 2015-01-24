package hello

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"log"
	"math/rand"
	"net/http"
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
		log.Println(len(keys))

		// Return a random gifID for now
		gifId := keys[rand.Intn(len(keys))]
		gif := gifs[gifId]
		content := gif.content
		r.JSON(200, map[string]interface{}{
			"id":      gifId,
			"title":   gif.gifTitle,
			"game":    gif.gameTitle,
			"upvotes": content.upvotes,
		})
	})

	m.Get("/scrape", func(res http.ResponseWriter, req *http.Request) {
		gifs := scrapeSubreddit("smashbros", req)
		keys = extractKeys(gifs)
	})

	http.Handle("/", m)
}
