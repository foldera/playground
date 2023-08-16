package release1_21

import (
	"cmp"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRelease121_Min(t *testing.T) {
	type testCase[T cmp.Ordered] struct {
		name           string
		inputs         []T
		expectedMinMax [2]T
	}

	testCases := []testCase[int]{
		{"min&max(1,2)", []int{1, 2}, [2]int{1, 2}},
		{"min&max(-1,2)", []int{-1, 2}, [2]int{-1, 2}},
		{"min&max(-10)", []int{-10}, [2]int{-10, -10}},
		{"min&max(-10,-20,53)", []int{-10, -20, 53}, [2]int{-20, 53}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			switch len(tc.inputs) {
			case 1:
				assert.Equal(t, tc.expectedMinMax[0], min(tc.inputs[0]))
				assert.Equal(t, tc.expectedMinMax[1], max(tc.inputs[0]))
			case 2:
				assert.Equal(t, tc.expectedMinMax[0], min(tc.inputs[0], tc.inputs[1]))
				assert.Equal(t, tc.expectedMinMax[1], max(tc.inputs[0], tc.inputs[1]))
			case 3:
				assert.Equal(t, tc.expectedMinMax[0], min(tc.inputs[0], tc.inputs[1], tc.inputs[2]))
				assert.Equal(t, tc.expectedMinMax[1], max(tc.inputs[0], tc.inputs[1], tc.inputs[2]))
			default:
				assert.True(t, true)
			}

		})
	}

}
