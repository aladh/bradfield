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

	done <- struct{}{} // swap this with the one inside the goroutine, no need to make it unbuffered
	fmt.Println("initialization done, continuing with rest of program")
}
