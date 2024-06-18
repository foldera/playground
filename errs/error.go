package errs

import (
	"fmt"
)

type Code interface {
	Index() int
	Err() error
}

type Error interface {
	error
	Code() Code
	CausedBy(parent error) error
}

func New(code Code, msg string) Error {
	if code == nil {
		return nil
	}
	var err error
	if msg != "" {
		err = fmt.Errorf("[%d][%w]::%s", code.Index(), code.Err(), msg)
	} else {
		err = fmt.Errorf("[%d][%w]", code.Index(), code.Err())
	}
	return &custom{msg: msg, code: code, error: err}
}

type Handler func(err error) Error

func Handle(err error, handlers ...Handler) error {
	switch err.(type) {
	case nil:
		return nil
	case Error:
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
