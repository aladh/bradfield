package main

import (
	"fmt"
	"log"
)

func main() {
	var command string
	var arg string

	for {
		fmt.Print("Enter a command: ")

		_, err := fmt.Scanf("%s %s", &command, &arg)
		if err != nil {
			log.Fatalf("error reading input: %s\n", err)
		}

		fmt.Println(command, arg)
	}
}
