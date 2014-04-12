package api

import (
	"github.com/go-martini/martini"
)

func addNeedRoutes(r martini.Router) {
	r.Group("/needs", func(r martini.Router) {
		r.Get("", func() string {
			return "All my needs"
		})
		r.Get("/:id", func(p martini.Params) string {
			return "Need(" + p["id"] + ")"
		})
	})
}
