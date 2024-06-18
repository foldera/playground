package errs_test

import (
	"errors"
)

var (
	internalServerError = errors.New("internal server")
	unauthorizedError   = errors.New("unauthorized")
	forbiddenError      = errors.New("forbidden")
	validationError     = errors.New("validation")
	notFoundError       = errors.New("not found")
	duplicationError    = errors.New("duplication")
)

const (
	InternalServer code = iota
	Unauthorized
	Forbidden
	Validation
	NotFound
	Duplication
)

type code uint8

func (c code) Index() int {
	return int(c)
}

func (c code) Err() error {
	switch c {
	case Unauthorized: // 1
		return unauthorizedError
	case Forbidden: // 2
		return forbiddenError
	case Validation: // 3
		return validationError
	case NotFound: // 4
		return notFoundError
	case Duplication: // 5
		return duplicationError
	default: // 0
		return internalServerError
	}
}
