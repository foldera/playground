package pool

import (
	"sync"
	"testing"
)

func BenchmarkSyncPool_BuiltinMakeSlice(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		wg := &sync.WaitGroup{}
		for i := 0; i < maxGoroutineNumbers; i++ {
			wg.Add(1)
			go func(wg *sync.WaitGroup) {
				defer wg.Done()
				s := make([]int, 0, sliceCapacity)
				fill(s, sliceCapacity, 1237)
			}(wg)
		}
		wg.Wait()
	}
}

func BenchmarkSyncPool_SlicePool(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		wg := &sync.WaitGroup{}
		for i := 0; i < maxGoroutineNumbers; i++ {
			wg.Add(1)
			go func(wg *sync.WaitGroup) {
				defer wg.Done()
				s := slicePool.Get().([]int)
				fill(s, sliceCapacity, 1)
				s = s[:0]
				slicePool.Put(s)
			}(wg)
		}
		wg.Wait()
	}
}
