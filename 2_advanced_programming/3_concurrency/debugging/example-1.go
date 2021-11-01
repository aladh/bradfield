package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		// Copy i into goroutine to avoid shared access
		go func(n int) {
			fmt.Printf("launched goroutine %d\n", n)
		}(i)
	}
	// Wait for goroutines to finish
	time.Sleep(time.Second)
}
