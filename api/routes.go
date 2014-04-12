package api

import (
	"net/http"

	"github.com/go-martini/martini"
)

func Must(data string, err error) string {
	if err != nil {
		panic(err)
	}
	return data
}

type Encoder interface {
	Encode(v interface{}) (string, error)
}

type Decoder interface {
	Decode(d string, v interface{}) error
}

type jsonEncoder struct{}

type jsonDecoder struct{}

func (_ jsonDecoder) Decode(d string, v interface{}) error {
	return nil
}

func MapCodec(c martini.Context, r *http.Request) {
	if r.Method == "POST" || r.Method == "PUT" {
		c_type := r.Header["Content-Type"]
		if len(c_type) != 1 {
			panic("That's unreasonable")
		}
		encoder := map[string]Decoder{
			"application/json": jsonDecoder{},
		}[c_type[0]]

		if encoder == nil {
		}
	} else {
		c.MapTo(jsonEncoder{}, (*Encoder)(nil))
	}
}

func AddRoutes(r martini.Router) {
	r.Group("/api/v1", func(r martini.Router) {
		addNeedRoutes(r)
		addResourceRoutes(r)
		addStaffRoutes(r)
	}, MapCodec)
}
