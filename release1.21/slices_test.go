package release1_21

import (
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
)

// https://pkg.go.dev/slices

func TestRelease121_Slices_Contains(t *testing.T) {
	numbers := []int{1, 2, 3, 3, 4, 5}
	assert.True(t, slices.Contains(numbers, 3))
	assert.True(t, slices.Contains(numbers, 5))
	assert.False(t, slices.Contains(numbers, -1))
	assert.False(t, slices.Contains(nil, nil))
}
