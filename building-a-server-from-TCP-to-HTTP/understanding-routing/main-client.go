package understandingrouting

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// make a connection to the server
	conn, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	// write to the server
	// when connection is established we send messages to the server
	fmt.Fprintf(conn, "GET /index.html``\n")

	// read a response from the server
	bs := make([]byte, 1024)
	n, err := conn.Read(bs)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bs[:n]))
}
