package api

import (
	"github.com/go-martini/martini"
)

func addNeedRoutes(r martini.Router) {
	r.Get("/:id", func(p martini.Params) string {
		return "Nothing to see here: " + p["id"]
	})
}
