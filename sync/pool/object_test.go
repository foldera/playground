package pool

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

var maxGoroutineNumbers = 1000

func TestSyncPool_Object(t *testing.T) {
	modify := objectModifier()
	assert.Equal(t, new(Object), objectPool.Get().(*Object))

	first := new(Object)
	modify(first)
	objectPool.Put(first)
	assert.Equal(t, first, objectPool.Get().(*Object))

	second := new(Object)
	modify(second)
	objectPool.Put(second)
	assert.Equal(t, second, objectPool.Get().(*Object))

	second.Title = "second(edited)"
	objectPool.Put(second)
	assert.Equal(t, second, objectPool.Get().(*Object))
	assert.Equal(t, new(Object), objectPool.Get().(*Object))
	assert.Equal(t, new(Object), objectPool.Get().(*Object))
	assert.Equal(t, new(Object), objectPool.Get().(*Object))

	objectPool.Put(second)
	assert.Equal(t, second, objectPool.Get().(*Object))
	assert.Equal(t, new(Object), objectPool.Get().(*Object))
}

func BenchmarkSyncPool_BuiltinNewObject(b *testing.B) {
	b.ReportAllocs()
	change := objectModifier()
	for n := 0; n < b.N; n++ {
		wg := &sync.WaitGroup{}
		for i := 0; i < maxGoroutineNumbers; i++ {
			wg.Add(1)
			go func(wg *sync.WaitGroup) {
				defer wg.Done()
				object := new(Object)
				change(object)
			}(wg)
		}
		wg.Wait()
	}
}

func BenchmarkSyncPoolObject_ObjectPool(b *testing.B) {
	b.ReportAllocs()
	change := objectModifier()
	for n := 0; n < b.N; n++ {
		wg := &sync.WaitGroup{}
		for i := 0; i < maxGoroutineNumbers; i++ {
			wg.Add(1)
			go func(wg *sync.WaitGroup) {
				defer wg.Done()
				object := objectPool.Get().(*Object)
				change(object)
				object.Id = 0
				object.Title = ""
				slicePool.Put(object)
			}(wg)
		}
		wg.Wait()
	}
}
