package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// underhood implementation of http
func main() {
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
}
