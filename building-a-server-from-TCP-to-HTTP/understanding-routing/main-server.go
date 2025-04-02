package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	// server sits and wait at door 8080
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	defer ln.Close()

	for {
		/*
			Accept return Connection and connection returns
			Write - writes data to the connection (can be limited by time and
			return error)
			Read - reads data from the connection, can be made to time out and return
			error after time limit - same as write
		*/
		// when client come in into 8080
		// we accept connection
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}

		go handleConnection(conn)
		/*
			using just handleConnection will block the flow
			and one message would be processed at the time.
			Because of that is very important to
			use goroutine to handle multiple connections
		*/

	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// create a new Reader
	reader := bufio.NewReader(conn)

	// read command line from the client
	line, err := reader.ReadString()

	if err != nil {
		fmt.Fprintf(conn, "error reading command %v\n", err)
		return
	}

	// Routing
	/*
		The idea is that the request comes in and we handle that request
		or we sent that request to the right handler.
		That is the whole idea that we're going to be doing with the HTTP.
	*/

	parts := strings.SplitN(strings.TrimSpace(line), "/", 2)
	if len(parts) != 2 {
		fmt.Fprintf(conn, "Invalid command format. Expected format: COMMAND:RESOURCE\n")
		return
	}

	// recive message from the client
	command := parts[0]
	resource := parts[1]

	log.Printf("Receiver command: %s %s\n", command, resource)

	// handle recived message in the proper way
	switch command {
	case "GET":
		handleGet(conn, resource)
	default:
		fmt.Fprintf(conn, "unknown command : %s\n", command)
	}
}

func handleGet(conn net.Conn, resource string) {
	fmt.Fprintf(conn, "recived message is %s", resource)
}
