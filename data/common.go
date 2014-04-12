package data

import (
	"time"

	"appengine"
	"appengine/datastore"
	"appengine/user"
)

type commonBase struct {
	Who         string              `db:"owner"`
	Creation    time.Time           `db:""`
	What        []string            `db:"m=0:o+w,m=1:o+a"`
	Attachments []appengine.BlobKey `db:"m=0:o+w,m=1:o+a"`
	Categories  []string            `db:"m=0:o+w"`
	Expiration  time.Time           `db:"m=0:o+w"`

	AdminCategories []string       `db:"m=0|1:a+w"`
	Match           *datastore.Key `db:"m=0|1:a+w,o-r"`
	Complete        bool           `db:"m=0|1:a+w"`
}

func newCommonBase(c appengine.Context) *commonBase {
	return &commonBase{
		Who:      user.Current(c).String(),
		Creation: time.Now(),
	}
}
