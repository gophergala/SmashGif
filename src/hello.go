package hello

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
)

var gifs map[string]Gif

func init() {
	m := martini.Classic()
	// Middle ware stuff
	m.Use(render.Renderer(render.Options{
		Directory: "public",
	}))
	m.Use(martini.Logger())
	m.Use(martini.Static("public"))

	// Request handlers
	m.Get("/", func(r render.Render, req *http.Request) {
		r.HTML(200, "main", nil)
	})

	m.Get("/api", func(r render.Render, req *http.Request) {
		qs := req.URL.Query()
		r.JSON(200, map[string]interface{}{"id": qs})
	})

	http.Handle("/", m)
}

func main() {
	gifs = scrapeSubreddit("smashbros")
}
