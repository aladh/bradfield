package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
)

const prompt = "⛄  "
const exitMessage = "❄❅❄❅ Goodbye and stay warm! ❄❅❄❅"
const inputSeparator = " "
const exitCommand = "exit"

var interruptChan = make(chan os.Signal, 1)

func main() {
	signal.Notify(interruptChan, os.Interrupt)

	var command string
	flag.StringVar(&command, "c", "", "Run a command and exit")
	flag.Parse()

	if len(command) > 0 {
		runCommand(command)
		os.Exit(0)
	}

	scanner := bufio.NewScanner(os.Stdin)

CommandLoop:
	for {
		fmt.Print(prompt)

		if ok := scanner.Scan(); !ok {
			break
		}

		input := scanner.Text()

		switch input {
		case exitCommand:
			break CommandLoop
		default:
			runCommand(input)
		}
	}

	fmt.Println(exitMessage)
}

func runCommand(input string) {
	splitCommand := strings.Split(input, inputSeparator)
	commandName := splitCommand[0]
	commandPath, err := exec.LookPath(commandName)
	if err != nil {
		fmt.Printf("error finding command: %s\n", err)
		return
	}

	attrs := os.ProcAttr{
		Dir:   "",
		Env:   nil,
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		Sys:   nil,
	}

	process, err := os.StartProcess(commandPath, splitCommand, &attrs)
	if err != nil {
		fmt.Printf("error running subprocess: %s\n", err)
	}

	doneChan := make(chan bool, 1)
	go forwardInterrupt(process, doneChan)

	_, err = process.Wait()
	if err != nil {
		fmt.Printf("error waiting for subprocess: %s\n", err)
	}

	doneChan <- true
}

func forwardInterrupt(process *os.Process, doneChan <-chan bool) {
	select {
	case <-interruptChan:
		err := process.Signal(os.Interrupt)
		if err != nil {
			fmt.Printf("error signalling subprocess: %s\n", err)
		}
		fmt.Println() // Add newline after ^C
	case <-doneChan:
		break
	}
}
