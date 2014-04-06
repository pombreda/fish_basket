package fish_basket

import (
	"net/http"

	"github.com/go-martini/martini"

	"github.com/riannucci/fish_basket/api"
)

func makeDirtyMartini() (*martini.Martini, martini.Router) {
	r := martini.NewRouter()
	m := martini.New()

	m.Use(martini.Logger())
	m.Use(martini.Recovery())
	m.MapTo(r, (*martini.Routes)(nil))
	m.Action(r.Handle)

	return m, r
}

func Run() {
	m, r := makeDirtyMartini();
	api.AddRoutes(r)

	r.Get("/", func() string {
		return "Hello bob!"
	})

	http.Handle("/", m)
}
