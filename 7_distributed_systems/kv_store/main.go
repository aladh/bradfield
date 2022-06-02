package main

import (
	"fmt"
	"log"
	"strings"
)

const GetCommand = "get"
const SetCommand = "set"

func main() {
	data := make(map[string]string)
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

		switch command {
		case GetCommand:
			fmt.Println(data[arg])
		case SetCommand:
			splitArg := strings.Split(arg, "=")
			data[splitArg[0]] = splitArg[1]
		}
	}
}
