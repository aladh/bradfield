package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"

	"kv_store/commands"
)

const sockName = "server.sock"

func main() {
	var command string
	var arg string

	sockAddr, err := SockAddr()
	if err != nil {
		log.Fatalf("error getting socket address: %s\n", err)
	}

	conn, err := net.Dial("unix", sockAddr)
	if err != nil {
		log.Fatalf("error connecting to server: %s\n", err)
	}

	for {
		fmt.Print("Enter a command: ")

		_, err := fmt.Scanf("%s %s", &command, &arg)
		if err != nil {
			log.Fatalf("error reading input: %s\n", err)
		}

		if command != commands.GetCommand && command != commands.SetCommand {
			fmt.Printf("invalid command: %s\n", command)
			continue
		}

		message := fmt.Sprintf("%s %s\n", command, arg)
		_, err = conn.Write([]byte(message))
		if err != nil {
			log.Fatalf("error writing to connection: %s\n", err)
		}

		reader := bufio.NewReader(conn)
		resp, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("error reading from connection: %s\n", err)
		}

		fmt.Print(resp)
	}
}

func SockAddr() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("error getting working directory: %s", err)
	}

	return fmt.Sprintf("%s/%s", wd, sockName), nil
}
