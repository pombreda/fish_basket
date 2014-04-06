package api

import (
	"github.com/go-martini/martini"
)

func AddRoutes(r martini.Router) {
	r.Group("/api", func(r martini.Router) {
		addNeedRoutes(r)
	})
}

