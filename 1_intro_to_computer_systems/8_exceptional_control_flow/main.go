package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

const prompt = "⛄  "
const exitMessage = "❄❅❄❅ Goodbye and stay warm! ❄❅❄❅"

func main() {
	var command string
	flag.StringVar(&command, "c", "", "Run a command and exit")
	flag.Parse()

	if len(command) > 0 {
		runCommand(command)
		os.Exit(0)
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(prompt)

		if ok := scanner.Scan(); !ok {
			fmt.Println(exitMessage)
			break
		}

		runCommand(scanner.Text())
	}
}

func runCommand(command string) {
	fmt.Println(command)
}
