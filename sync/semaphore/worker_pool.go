package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"log"
	"time"
)

var (
	nJobs    = 100
	nWorkers = 4
	sem      = semaphore.NewWeighted(int64(nWorkers))
)

func main() {
	log.Printf("start to calculate the sqaure of {0,1,2,.....,%d} using %d workers", nJobs-1, nWorkers)

	results := make([]int, nJobs)
	ctx := context.TODO()
	for i := range results {
		if err := sem.Acquire(ctx, 1); err != nil {
			fmt.Println("cannot acquire semaphore:", err)
			break
		}
		go func(n int) {
			defer sem.Release(1)
			results[n] = square(n)
		}(i)
	}
	err := sem.Acquire(ctx, int64(nWorkers))
	if err != nil {
		fmt.Println(err)
	}
	for n, nSquare := range results {
		fmt.Println(n, "->", nSquare)
	}
	log.Println("done!")
}

// square simulates a worker which calculates the square of a given integer(task id), sleeps 1 second, and returns the result.
func square(n int) int {
	s := n * n
	time.Sleep(time.Second)
	return s
}
