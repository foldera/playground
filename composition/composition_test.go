package composition

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComposition(t *testing.T) {
	a := ObjectA{1, "A", 5}
	b := ObjectB{2, "B", "BBB"}
	ambiguous := AmbiguousObject{a, b, nil}
	assert.Equal(t, uint(1), ambiguous.ObjectA.ID)
	assert.Equal(t, uint(2), ambiguous.ObjectB.ID)
	//assert.Equal(t, ..., ambiguous.ID)    // >>>>>>>>>>>> ambiguous.ID: ambiguous reference 'ID'
	//assert.Equal(t, ..., ambiguous.Title) // >>>>>>>>>>>> ambiguous.Title: ambiguous reference 'Title'
	//assert.Equal(t, ..., ambiguous.Count) // >>>>>>>>>>>> ambiguous.Count: ambiguous reference 'Count'
	assert.Equal(t, "BBB", ambiguous.Description)
	assert.Nil(t, ambiguous.ObjectC)

	// ##################################################################################
	c := ObjectC{"C", 3}
	clearObject := ClearObject{a, b, &c}

	assert.Equal(t, uint(1), clearObject.ID)
	assert.Equal(t, "A", clearObject.Title)
	assert.Equal(t, uint32(5), clearObject.Count)

	assert.Equal(t, uint(2), clearObject.B.ID)
	assert.Equal(t, "B", clearObject.B.Title)
	if assert.NotNil(t, clearObject.C) {
		assert.Equal(t, "C", clearObject.C.Value)
		assert.Equal(t, uint64(3), clearObject.C.Count)
	}

}
