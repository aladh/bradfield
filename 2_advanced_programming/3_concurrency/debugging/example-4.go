package main

import (
	"fmt"
)

func main() {
	// Make the channel unbuffered so sending blocks the main goroutine until the other goroutine is finished
	done := make(chan struct{})
	go func() {
		fmt.Println("performing initialization...")
		<-done
	}()

	done <- struct{}{}
	fmt.Println("initialization done, continuing with rest of program")
}
