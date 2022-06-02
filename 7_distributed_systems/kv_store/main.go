package main

import (
	"fmt"
	"log"
)

const GetCommand = "get"
const SetCommand = "set"

func main() {
	var command string
	var arg string

	for {
		fmt.Print("Enter a command: ")

		_, err := fmt.Scanf("%s %s", &command, &arg)
		if err != nil {
			log.Fatalf("error reading input: %s\n", err)
		}

		if command != GetCommand && command != SetCommand {
			fmt.Printf("invalid command: %s\n", command)
			continue
		}

		fmt.Println(command, arg)
	}
}
