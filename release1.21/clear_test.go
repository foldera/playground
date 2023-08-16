package release1_21

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRelease121_Clear(t *testing.T) {
	//	###################################################
	stringSlice := []string{"a", "b", "c"}
	clear(stringSlice)
	assert.Equal(t, []string{"", "", ""}, stringSlice)
	//	###################################################
	intSlice := []int{1, 2, 3}
	clear(intSlice)
	assert.Equal(t, []int{0, 0, 0}, intSlice)
	//	###################################################
	type testObject struct {
		id    uint
		title string
		is    bool
	}
	anySlice := []any{1, 2, 3.5, "string", testObject{2, "title", true}}
	clear(anySlice)
	assert.Equal(t, []any{nil, nil, nil, nil, nil}, anySlice)
	//	###################################################
	objectSlice := []testObject{{1, "first", true}, {2, "second", false}}
	clear(objectSlice)
	assert.Equal(t, []testObject{{0, "", false}, {0, "", false}}, objectSlice)
	//	###################################################
	map1 := map[string]int{"a": 1, "b": 2}
	clear(map1)
	assert.Equal(t, make(map[string]int, 0), map1)
	map1["c"] = 3
	assert.Equal(t, 3, map1["c"])
	clear(map1)
	assert.Equal(t, make(map[string]int, 0), map1)
	map1 = nil
	clear(map1)
	assert.Equal(t, map[string]int(nil), map1)
	//	###################################################
	map2 := map[uint]testObject{1: {1, "first", false}, 2: {2, "second", true}}
	clear(map2)
	assert.Equal(t, make(map[uint]testObject, 0), map2)
	//	###################################################
	//	###################################################
	//	###################################################

}
