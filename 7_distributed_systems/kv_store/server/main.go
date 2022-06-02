package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"kv_store/commands"
	"kv_store/kvdata"
)

const sockName = "server.sock"

func main() {
	kv, err := kvdata.Initialize()
	if err != nil {
		log.Fatalf("error initializing data: %s\n", err)
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("error getting working directory: %s\n", err)
	}

	sockAddr := fmt.Sprintf("%s/%s", wd, sockName)
	l, err := net.Listen("unix", sockAddr)
	defer l.Close()
	if err != nil {
		log.Fatalf("error listening on socket: %s\n", err)
	}

	log.Printf("Listening on socket %s\n", sockAddr)

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("error accepting connection: %s\n", err)
			continue
		}

		go handleConn(conn, kv)
	}
}

func handleConn(conn net.Conn, kv *kvdata.KVData) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message := scanner.Text()
		log.Printf("received message: %s\n", message)

		var command string
		var arg string

		_, err := fmt.Sscanf(message, "%s %s", &command, &arg)
		if err != nil {
			log.Fatalf("error reading message: %s\n", err)
		}

		switch command {
		case commands.GetCommand:
			val := kv.Get(arg)
			resp := fmt.Sprintf("%s\n", val)

			_, err = conn.Write([]byte(resp))
			if err != nil {
				log.Printf("error writing to connection: %s\n", err)
			}
		case commands.SetCommand:
			splitArg := strings.Split(arg, "=")
			resp := "OK\n"

			err := kv.Set(splitArg[0], splitArg[1])
			if err != nil {
				e := fmt.Sprintf("error setting value: %s\n", err)
				log.Print(e)
				resp = e
			}

			_, err = conn.Write([]byte(resp))
			if err != nil {
				log.Printf("error writing to connection: %s\n", err)
			}
		}
	}

	conn.Close()
}
