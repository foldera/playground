package pool

import "sync"

var (
	//	pool = sync.Pool{New: func() any {
	//		return &Object{1, "first"}
	//	}}
	pool = sync.Pool{New: func() any {
		return new(Object)
	}}
)

type Object struct {
	Id    uint
	Title string
}
