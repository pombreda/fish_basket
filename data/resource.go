package data

import "appengine"

type ResourceStatus int32

const (
	ResourcePending ResourceStatus = iota
	ResourceMatched
	ResourceExhausted
)

type Resource struct {
	*commonBase
}

func NewResource(c appengine.Context) *Resource {
	return &Resource{newCommonBase(c)}
}

func (r Resource) Status() ResourceStatus {
	if r.Complete {
		return ResourceExhausted
	}
	if r.Match != nil {
		return ResourceMatched
	}
	return ResourcePending
}

func (r Resource) DBAclMode() interface{} {
	return r.Status()
}
