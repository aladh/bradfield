package main

import (
	"fmt"
)

const numTasks = 3

func main() {
	// nil channel - need to instantiate it, no need to buffer
	// Buffer channel so goroutines aren't blocked
	done := make(chan struct{}, numTasks)
	for i := 0; i < numTasks; i++ {
		go func() {
			fmt.Println("running task...")

			// Signal that task is done
			done <- struct{}{}
		}()
	}

	// Wait for tasks to complete
	for i := 0; i < numTasks; i++ {
		<-done
	}
	fmt.Printf("all %d tasks done!\n", numTasks)
}
