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

func (x *custom) Unwrap() error {
	switch {
	case x == nil:
		return nil
	case errors.Unwrap(x.error) != x.code.Err():
		// x.CausedBy() called with non-nil parent error...
		return errors.Join(x.code.Err(), x.error)
	default:
		// New() called without calling CausedBy(), or called CausedBy(nil)
		return x.error
	}
}

func (x *custom) Trace() []error {
	if x == nil {
		return nil
	}
	result := []error{x.code.Err()}
	next := errors.Unwrap(x.error)
	if next == x.code.Err() {
		return result
	}
	for next != nil {
		result = append(result, next)
		next = errors.Unwrap(next)
	}
	return result
}

func (x *custom) CausedBy(parent error) error {
	if parent == nil {
		return x
	}
	x.error = fmt.Errorf("%v caused by (%w)", x.error, parent)
	return x
}

//func (x *custom) Is(target error) bool {
//	if errors.Is(x.code.Err(), target) {
//		return true
//	}
//	//if custom, ok := target.(*custom); ok {
//	//	return custom.code == e.code && custom.msg == e.msg
//	//}
//	next := errors.Unwrap(x.error)
//	for next != nil {
//		if errors.Is(next, target) {
//			return true
//		}
//		next = errors.Unwrap(next)
//	}
//	return false
//}
