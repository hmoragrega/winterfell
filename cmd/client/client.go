package main

// Client

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {

	address := flag.String("address", "127.0.0.1:8100", "Server address. Example: 127.0.0.1:8100")
	flag.Parse()

	conn, err := net.Dial("tcp", *address)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Connected to game server")
	defer conn.Close()

	stop := make(chan struct{})

	go printServerMessages(conn, stop)
	go sendPlayerCommands(conn)

	<-stop
}

func sendPlayerCommands(conn net.Conn) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		command := scanner.Text()

		if command == "" {
			continue
		}

		if _, err := conn.Write([]byte(command)); err != nil {
			log.Printf("Error sending the command: %s", err)
		}
	}
}

func printServerMessages(conn net.Conn, stop chan struct{}) {
	for {
		message, err := readServerMessage(conn)
		if err != nil && err == io.EOF {
			log.Printf("Server closed connection; terminating\n")
			break
		} else if err != nil {
			log.Printf("Error reading server message: %s\n", err)
		} else {
			fmt.Print(message)
		}
	}

	stop <- struct{}{}
}

func readServerMessage(conn net.Conn) (string, error) {
	var buffer = make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		return "", err
	}

	message := string(buffer[:])
	if message == "" {
		return "", fmt.Errorf("Read an empty command")
	}

	return message, nil
}
