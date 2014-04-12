package api

import (
	"github.com/go-martini/martini"
)

func addResourceRoutes(r martini.Router) {
	r.Get("/resources", func() string {
		return "All my resources"
	})

	r.Get("/resources/:id", func(p martini.Params) string {
		return "Resources(" + p["id"] + ")"
	})
}
