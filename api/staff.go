package api

import (
	"appengine"
	"appengine/user"

	"github.com/go-martini/martini"
)

func ensureStaffMember(c appengine.Context) {
	// TODO(riannucci): Restrict to actual staff members
	c.Infof("user accessing /staff is: %v", user.Current(c))
}

func addStaffRoutes(r martini.Router) {
	r.Group("/staff", func(r martini.Router) {
		r.Get("", func() string {
			return "Staff stuff?"
		})
	}, ensureStaffMember)
}
