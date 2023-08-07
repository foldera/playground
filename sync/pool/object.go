package pool

import (
	"fmt"
	"sync"
)

var objectPool = sync.Pool{New: func() any {
	return &Object{0, ""}
}}

type Object struct {
	Id    int
	Title string
}

func objectModifier() func(object *Object) {
	i := 1
	return func(object *Object) {
		object.Id = i
		object.Title = fmt.Sprintf("object(%d)", i)
		i++
	}
}
