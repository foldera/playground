package object

import (
	"github.com/foldera/playground/pointers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClosure_NewObject(t *testing.T) {
	type testCase struct {
		name     string
		params   Params
		expected *model
	}

	testCases := []testCase{
		{"with empty params", Params{}, new(model)},
		{"with id", Params{Id: pointers.New(uint(5))}, &model{id: 5}},
		{"with title", Params{Title: pointers.New("sample title")}, &model{title: "sample title"}},
		{"with id and title",
			Params{Id: pointers.New(uint(10)), Title: pointers.New("sample title")}, &model{id: 10, title: "sample title"},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			object := New()(tc.params)
			assert.NotNil(t, object)
			got, casted := object.(*model)
			if assert.True(t, casted) {
				assert.Equal(t, tc.expected, got)
			}
		})
	}
}
