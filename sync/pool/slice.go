/*
Golang has a lot of advantages such as rapid development/fast compilation and many more.
But for performance critical applications this language might not be the right language for you. why?
In order to add simplicity for the developer the language manages the memory for you. what does that mean?

Instead of managing the memory manually by allocating and deallocating memory in the heap when needed,
The Golang compiler decides during compilation which objects will be allocated on the heap and in runtime
has a program running in the background that scans the heap in intervals and deallocates memory that is not
being used.
The program that cleans the heap is called the Garbage Collector and can have a significant impact on performance.

solution(I): For performance-critical applications use manual memory-managed programming languages.

solution(II): Identify critical code bottlenecks and rewrite them in c and use cgo to run your
application but keep all other parts written in golang.

solution(III): Reuse memory in the heap using sync.Pool. Many programs create different instances
of the same object many times so instead of allocating new memory for each new object, we will have
a pool of these objects saved that we can reuse. That way we will not allocate new memory each time
we create this object and the Garbage collector will have less work to do. That's where sync.Pool comes in.
Ref: https://medium.com/@aryehlevklein/golang-using-sync-package-to-10x-performance-and-reduce-memory-footprint-a1ed4ee14931
*/

package pool

import "sync"

var sliceCapacity = 100000

var slicePool = sync.Pool{New: func() any {
	return make([]int, 0, sliceCapacity)
}}

func fill(s []int, maxLen int, value int) {
	if s == nil || cap(s) < maxLen {
		return
	}
	for i := 0; i < maxLen; i++ {
		s = append(s, value)
	}
}
