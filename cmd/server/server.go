package main

// Server

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"strings"

	"github.com/hmoragrega/winterfell/game"
)

func main() {
	address := flag.String("address", ":8100", "Server address to listen at. Example: :8100")
	flag.Parse()

	listen, error := net.Listen("tcp", *address)
	if error != nil {
		fmt.Println(error)
		return
	}

	fmt.Printf("Listening on address %s\n", *address)

	for {
		// TODO handle termination signals gracefully

		conn, err := listen.Accept()
		if err != nil {
			log.Printf("There has been an error accepting a connection: %s\n", err)
			continue
		}

		go handleNewConnection(conn)
	}
}

func executePlayerCommands(conn net.Conn, g game.Engine, stop chan struct{}) {
	for {
		command, parameters, err := readPlayerCommand(conn)
		if err != nil && err == io.EOF {
			log.Printf("Client disconnected\n")
			g.Stop()
			break
		} else if err != nil {
			log.Printf("Error reading player command: %s\n", err)
		} else {
			log.Printf("Received command %s\n", command)
			answer := executeGameCommand(g, command, parameters)
			if answer != "" {
				broadcastMessage(conn, answer)
			}
		}
	}

	stop <- struct{}{}
}

func brodcastGameMessages(conn net.Conn, g game.Engine, stop chan struct{}) {
	for {
		select {
		case <-stop:
			return
		case position := <-g.EnemyPosition():
			broadcastMessage(conn, getEnemyPositionMessage(g, position))
			break
		case result := <-g.GameOver():
			broadcastMessage(conn, getGameOverMessage(result))
			break
		}
	}
}

func handleNewConnection(conn net.Conn) {

	log.Println("Client connected")
	defer conn.Close()

	g := game.NewWinterIsComingEngine()
	stop := make(chan struct{})

	go executePlayerCommands(conn, g, stop)

	brodcastGameMessages(conn, g, stop)

	log.Println("Closing client connnection")
}

func broadcastMessage(conn net.Conn, message string) {
	log.Printf("Broadcasting message: %s", message)

	buffer := []byte(message)
	_, err := conn.Write(buffer)
	if err != nil {
		log.Printf("Error writting message to client %s\n", err)
	}
}

func readPlayerCommand(conn net.Conn) (string, []string, error) {
	var buffer = make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		return "", nil, err
	}

	message := cleanMessage(buffer)
	if message == "" {
		return "", nil, fmt.Errorf("Read an empty command")
	}

	log.Printf("Message length %d", len(message))

	parts := strings.Fields(message)

	return parts[0], parts[1:], nil
}

func cleanMessage(buffer []byte) string {
	buffer = bytes.Trim(buffer, "\x00")

	return string(buffer[:])
}
