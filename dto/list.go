package dto

import "fmt"

type List[T any] []T

func NewList[T any, L ~[]M, M any](models L, mappers ...Mapper[M, T]) (List[T], error) {
	var err error
	switch {
	case models == nil, len(models) == 0:
		var target T
		return nil, fmt.Errorf("trying to create new List[%T] using nil %T", target, models)

	case mappers == nil, len(mappers) == 0:
		var target T
		return nil, fmt.Errorf("trying to create new List[%T] without any mapper", target)

	default:
		list := make(List[T], len(models))
		for i := range models {
			for _, convert := range mappers {
				if list[i], err = convert(models[i]); err != nil {
					return nil, err
				}
			}
		}
		return list, nil
	}
}

func NewPtrList[T any, L ~[]M, M any](models L, mappers ...Mapper[M, T]) (List[*T], error) {
	list, err := NewList[T](models, mappers...)
	if err != nil {
		return nil, err
	}
	return NewList[*T](list, Ptr[T]())
}
