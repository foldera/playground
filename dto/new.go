package dto

import (
	"fmt"
	"reflect"
)

func New[T, F any](from F, convert Mapper[F, T]) (*T, error) {
	var (
		result T
		err    error
	)
	switch {
	case convert == nil:
		return nil, fmt.Errorf("trying to create new %T without any mapper", result)

	case reflect.ValueOf(from).IsZero():
		return nil, fmt.Errorf("trying to create new %T using zero value of %T", result, from)

	default:
		if result, err = convert(from); err != nil {
			return nil, err
		}
		return &result, nil
	}
}
