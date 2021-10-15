package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"path"
	"strings"
)

const prompt = "⛄  "
const exitMessage = "❄❅❄❅ Goodbye and stay warm! ❄❅❄❅"
const pathEnvVar = "PATH"
const pathSeparator = ":"
const inputSeparator = " "

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

func runCommand(input string) {
	splitCommand := strings.Split(input, inputSeparator)
	commandName := splitCommand[0]
	commandPath, err := findCommand(commandName)
	if err != nil {
		fmt.Println(err)
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

	_, err = process.Wait()
	if err != nil {
		fmt.Printf("error waiting for subprocess: %s\n", err)
	}
}

func findCommand(command string) (string, error) {
	pathLocations := strings.Split(os.Getenv(pathEnvVar), pathSeparator)
	for _, dir := range pathLocations {
		entries, err := os.ReadDir(dir)
		if err != nil {
			fmt.Printf("error reading PATH directory: %s\n", err)
		}

		for _, entry := range entries {
			if !entry.IsDir() && entry.Name() == command {
				return path.Join(dir, entry.Name()), nil
			}
		}
	}

	return "", errors.New(fmt.Sprintf("%s: command not found\n", command))
}
