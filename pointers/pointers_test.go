package pointers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPointers_New(t *testing.T) {
	type testCase[T SupportedType] struct {
		name  string
		input T
	}
	/*#####################################
	############################## booleans
	#######################################*/
	boolTestCases := []testCase[bool]{
		{"false boolean", false},
		{"true boolean", true},
	}
	for _, tc := range boolTestCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, &tc.input, New(tc.input))
		})
	}
	/*#####################################
	################################# uints
	#######################################*/
	uintTestCases := []testCase[uint]{
		{"zero uint", uint(0)},
		{"happy uint", uint(9)},
	}
	for _, tc := range uintTestCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, &tc.input, New(tc.input))
		})
	}
	/*#####################################
	############################### strings
	#######################################*/
	stringTestCases := []testCase[string]{
		{"empty string", ""},
		{"happy string", "happy"},
	}
	for _, tc := range stringTestCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, &tc.input, New(tc.input))
		})
	}
}

func TestPointers_Value(t *testing.T) {
	type testCase[T SupportedType] struct {
		name     string
		input    *T
		expected T
	}
	/*###################################
	########################################################
	############################## booleans
	#######################################*/
	boolTestCases := []testCase[bool]{
		{"nil boolean", nil, false},
		{"false boolean", New(false), false},
		{"true boolean", New(true), true},
	}
	for _, tc := range boolTestCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, Value(tc.input))
		})
	}
	/*#####################################
	################################# uints
	#######################################*/
	uintTestCases := []testCase[uint]{
		{"nil uint", nil, uint(0)},
		{"zero uint", New(uint(0)), uint(0)},
		{"happy uint", New(uint(9)), uint(9)},
	}
	for _, tc := range uintTestCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, Value(tc.input))
		})
	}
	/*#####################################
	############################### strings
	#######################################*/
	stringTestCases := []testCase[string]{
		{"nil string", nil, ""},
		{"empty string", New(""), ""},
		{"happy string", New("happy"), "happy"},
	}
	for _, tc := range stringTestCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, Value(tc.input))
		})
	}
}
