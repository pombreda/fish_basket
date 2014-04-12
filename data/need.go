package data

import "appengine"

type NeedStatus int32

const (
	NeedPending NeedStatus = iota
	NeedMatched
	NeedCompleted
)

type Need struct {
	*commonBase
}

func NewNeed(c appengine.Context) *Need {
	return &Need{newCommonBase(c)}
}

func (n Need) Status() NeedStatus {
	if n.Complete {
		return NeedCompleted
	}
	if n.Match != nil {
		return NeedMatched
	}
	return NeedPending
}

func (n Need) DBAclMode() int32 {
	return int32(n.Status())
}
