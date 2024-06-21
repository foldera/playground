package dto

import "fmt"

type List[T any] []T

func (l List[T]) Append(elems ...T) List[T] {
	return append(l, elems...)
}

func (l List[T]) Slice() []T {
	return l
}

func NewList[T any, L ~[]M, M any](models L, convert Mapper[M, T]) (List[T], error) {
	var err error
	switch {
	case convert == nil:
		var target T
		return nil, fmt.Errorf("trying to create new List[%T] without any mapper", target)

	case models == nil, len(models) == 0:
		var target T
		return nil, fmt.Errorf("trying to create new List[%T] using nil or empty %T", target, models)

	default:
		list := make(List[T], len(models))
		for i := range models {
			if list[i], err = convert(models[i]); err != nil {
				return nil, err
			}
		}
		return list, nil
	}
}

func NewPtrList[T any, L ~[]M, M any](models L, convert Mapper[M, T]) (List[*T], error) {
	list, err := NewList[T](models, convert)
	if err != nil {
		return nil, err
	}
	return NewList[*T](list, Ptr[T]())
}
