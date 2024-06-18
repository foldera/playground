package errs_test

import (
	"errors"
	"fmt"
	"github.com/foldera/playground/errs"
	"github.com/stretchr/testify/assert"
	"io/fs"
	"testing"
)

func TestNew(t *testing.T) {
	cause1 := errors.New("cause1")
	cause2 := fmt.Errorf("cause2 caused by %w", cause1)

	notFound := errs.New(NotFound, "something not found")
	assert.True(t, errors.Is(notFound, NotFound.Err()))
	assert.False(t, errors.Is(notFound, cause1))
	assert.False(t, errors.Is(notFound, cause2))

	notFoundCausedBy := errs.New(NotFound, "something not found").CausedBy(cause2)
	assert.True(t, errors.Is(notFoundCausedBy, cause1))
	assert.True(t, errors.Is(notFoundCausedBy, cause2))
	assert.True(t, errors.Is(notFoundCausedBy, NotFound.Err()))

	assert.False(t, errors.Is(notFound, notFoundCausedBy))

	var x errs.Error
	assert.True(t, errors.As(notFound, &x))
	assert.True(t, errors.As(notFoundCausedBy, &x))

}

func pathErrFirstHandler(path any) errs.Handler {
	return func(err error) errs.Error {
		if errors.Is(err, fs.ErrNotExist) {
			return errs.New(NotFound, fmt.Sprintf("%v not found", path))
		} else if errors.Is(err, fs.ErrExist) {
			return errs.New(Duplication, fmt.Sprintf("%v already exists", path))
		}
		return nil // Allowing next handler to be executed
	}
}
func pathErrFinisherHandler(path any) errs.Handler {
	return func(err error) errs.Error {
		if errors.Is(err, fs.ErrInvalid) {
			return errs.New(Validation, fmt.Sprintf("%v is not valid", path))
		} else if errors.Is(err, fs.ErrPermission) {
			return errs.New(Forbidden, fmt.Sprintf("%v is forbidden", path))
		}
		return errs.New(InternalServer, "something went wrong") // This is a finisher
	}
}

func TestHandle(t *testing.T) {
	pathHandlers := []errs.Handler{
		pathErrFirstHandler("path"),
		pathErrFinisherHandler("path"),
	}
	type expected struct {
		code errs.Code
		msg  string
	}
	testCases := []struct {
		name     string
		input    error
		handlers []errs.Handler
		expected *expected
	}{
		{"with nil input error", nil, nil, nil},
		{"with custom input error", errs.New(Forbidden, "msg"), nil, &expected{Forbidden, "msg"}},
		{"with nil handler", errors.New("any error"), nil, nil},
		{"without finisher", errors.New("any error"), []errs.Handler{pathErrFirstHandler("path")}, nil},
		{"path not exists", fs.ErrNotExist, pathHandlers, &expected{NotFound, "path not found"}},
		{"path exists", fs.ErrExist, pathHandlers, &expected{Duplication, "path already exists"}},
		{"invalid path", fs.ErrInvalid, pathHandlers, &expected{Validation, "path is not valid"}},
		{"forbidden path", fs.ErrPermission, pathHandlers, &expected{Forbidden, "path is forbidden"}},
		{"path closed", fs.ErrClosed, pathHandlers, &expected{InternalServer, "something went wrong"}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := errs.Handle(tc.input, tc.handlers...)
			// Ensure the errs.Handle function
			//1. will return nil in case of nil input error,
			//2. will return exactly the same error in case of any custom input,
			//3. will wrap any other error in the result error
			assert.True(t, errors.Is(err, tc.input))
			// Check the result error type
			var got errs.Error
			if errors.As(err, &got) {
				assert.Equal(t, tc.expected.code, got.Code())
				assert.Equal(t, tc.expected.msg, got.Error())
				assert.True(t, errors.Is(err, tc.input))
			}
		})
	}
}
