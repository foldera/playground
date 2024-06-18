package errs

import "fmt"

type Code interface {
	Index() int
	Err() error
}

type Reporter interface {
	error
	Code() Code
	Trace() []error
}

type Error interface {
	Reporter
	CausedBy(parent error) Reporter
}

func New(code Code, msg string) Error {
	if code == nil {
		return nil
	}
	return &custom{
		msg: msg, code: code,
		error: fmt.Errorf("[%d][%w]::%q", code.Index(), code.Err(), msg),
	}
}

type Handler func(err error) Error

func Handle(err error, handlers ...Handler) error {
	switch err.(type) {
	case nil, Error:
		return err
	default:
		for _, handle := range handlers {
			if handle != nil {
				if customErr := handle(err); customErr != nil {
					return handle(err).CausedBy(err)
				}
			}
		}
		return err
	}
}
