package main

import (
	"fmt"
	"log"
	"strings"

	"kv_store/kvdata"
)

const getCommand = "get"
const setCommand = "set"

func main() {
	var command string
	var arg string

	kv, err := kvdata.Initialize()
	if err != nil {
		log.Fatalf("error initializing data: %s\n", err)
	}

	for {
		fmt.Print("Enter a command: ")

		_, err := fmt.Scanf("%s %s", &command, &arg)
		if err != nil {
			log.Fatalf("error reading input: %s\n", err)
		}

		if command != getCommand && command != setCommand {
			fmt.Printf("invalid command: %s\n", command)
			continue
		}

		switch command {
		case getCommand:
			fmt.Println(kv.Get(arg))
		case setCommand:
			splitArg := strings.Split(arg, "=")
			err := kv.Set(splitArg[0], splitArg[1])
			if err != nil {
				log.Printf("error setting value: %s\n", err)
			}
		}
	}
}
