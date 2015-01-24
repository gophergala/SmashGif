package hello

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
)

func init() {
	m := martini.Classic()

	// Middle ware stuff
	m.Use(render.Renderer())
	m.Use(martini.Static("public"))

	// Request handlers
	m.Get("/", func(r render.Render) {
		r.HTML(200, "index", nil)
	})

	http.Handle("/", m)
}
