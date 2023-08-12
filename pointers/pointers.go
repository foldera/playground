package pointers

type SupportedType interface {
	~bool | ~uint | ~string
}

// New returns the pointer address of the given SupportedType value.
func New[T SupportedType](value T) *T {
	return &value
}

// Value returns the value of the given SupportedType pointer or its zero value if the pointer is nil.
func Value[T SupportedType](p *T) T {
	if p != nil {
		return *p
	}
	return *new(T)
}
