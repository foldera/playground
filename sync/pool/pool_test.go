package pool

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestPool(t *testing.T) {
	assert.Equal(t, new(Object), pool.Get().(*Object))
	first := &Object{1, "first"}
	pool.Put(first)
	assert.Equal(t, first, pool.Get().(*Object))
	second := &Object{2, "second"}
	pool.Put(second)
	assert.Equal(t, second, pool.Get().(*Object))
	second.Title = "second(edited)"
	pool.Put(second)
	assert.Equal(t, second, pool.Get().(*Object))
	assert.Equal(t, new(Object), pool.Get().(*Object))
	assert.Equal(t, new(Object), pool.Get().(*Object))
	assert.Equal(t, new(Object), pool.Get().(*Object))
	pool.Put(second)
	assert.Equal(t, second, pool.Get().(*Object))
	assert.Equal(t, new(Object), pool.Get().(*Object))
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		pool.Put(second)
	}()
	go func() {
		defer wg.Done()
		pool.New()
	}()
	wg.Wait()
	assert.Equal(t, second, pool.Get().(*Object))
}
