package main

import (
	"fmt"
	"sync"
)

type idService interface {
	// Returns values in ascending order; it should be safe to call
	// getNext() concurrently without any additional synchronization.
	getNext() uint64
}

func main() {
	test("noSyncService", &noSyncService{})         // race + sometimes not monotonic + final value too small
	test("atomicService", &atomicService{})         // no race
	test("mutexService", &mutexService{})           // no race
	test("goroutineService", newGoroutineService()) // no race
}

func test(name string, service idService) {
	wg := sync.WaitGroup{}

	for i := 0; i < 4; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			last := uint64(0)

			for j := 0; j < 1_000; j++ {
				val := service.getNext()

				if val == last {
					fmt.Printf("worker %d: last value %d matches returned value\n", id, last)
				} else {
					last = val
				}
			}
		}(i)
	}

	wg.Wait()
	fmt.Printf("%s final value: %d\n", name, service.getNext())
}
