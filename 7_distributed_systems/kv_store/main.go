package main

import (
	"fmt"
	"log"
	"strings"

	"kv_store/kvdata"
)

const GetCommand = "get"
const SetCommand = "set"

func main() {
	var command string
	var arg string

	kv := kvdata.LoadOrInitialize()

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
			fmt.Println(kv.Get(arg))
		case SetCommand:
			splitArg := strings.Split(arg, "=")
			kv.Set(splitArg[0], splitArg[1])
		}
	}
}
