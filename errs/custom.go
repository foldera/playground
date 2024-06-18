package errs

import (
	"errors"
	"fmt"
)

type custom struct {
	msg  string // wrapper message
	code Code   // your internal Code
	error
}

func (x *custom) Code() Code    { return x.code }
func (x *custom) Error() string { return x.msg }
func (x *custom) CausedBy(parent error) error {
	if parent == nil {
		return x
	}
	x.error = fmt.Errorf("%v caused by (%w)", x.error, parent)
	return x
}

func (x *custom) Is(target error) bool {
	if errors.Is(x.code.Err(), target) {
		return true
	}
	//if custom, ok := target.(*custom); ok {
	//	return custom.code == e.code && custom.msg == e.msg
	//}
	next := errors.Unwrap(x.error)
	for next != nil {
		if errors.Is(next, target) {
			return true
		}
		next = errors.Unwrap(next)
	}
	return false
}
