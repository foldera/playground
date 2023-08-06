package pool

import "sync"

var defaultObject = &Object{0, "default"}

var (
	pool = sync.Pool{New: func() any {
		return defaultObject
	}}
)

type Object struct {
	Id    uint
	Title string
}
