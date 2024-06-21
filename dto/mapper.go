package dto

type (
	Mapper[T1, T2 any] func(from T1) (T2, error)
)

func Ptr[T any]() Mapper[T, *T] {
	return func(from T) (*T, error) {
		return &from, nil
	}
}

func Value[T any]() Mapper[*T, T] {
	return func(from *T) (T, error) {
		var zero T
		if from == nil {
			return zero, nil
		}
		return *from, nil
	}
}
