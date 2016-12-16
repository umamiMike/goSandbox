package main

import (
	"fmt"
	"io"
	"net"
	"os"

	"fknsrs.biz/p/mllp"
	"github.com/facebookgo/stackerr"
)

const (
	ConnHost = ""
	ConnPort = "2575"
	ConnType = "tcp"
)

func main() {
	// Listen for incoming connections.
	l, err := net.Listen(ConnType, ":"+ConnPort)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + ConnHost + ":" + ConnPort)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		//logs an incoming message
		fmt.Printf("Received message %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())

		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	defer conn.Close()
	reader := mllp.NewReader(conn)
	for {
		message, err := reader.ReadMessage()
		if err != nil {
			if stackerr.HasUnderlying(err, stackerr.Equals(io.EOF)) {
				break
			}

			panic(err)
		} else {
			fmt.Println(string(message))
		}
		conn.Write(message)
	}
	conn.Close()
}
