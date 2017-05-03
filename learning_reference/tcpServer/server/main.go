package main

import (
	"bufio"
	"fmt"
	"net"
	"strings" // only needed below for sample processing
)

func main() {
	ln, err = net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Errorf("error in connecting")
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Errorf("error in accepting connection")
		}
		go handleConnection(conn)
	}
	fmt.Sprintln(ln, "the linster is of type: %T")
	// accept connection on port

}

func handleConnection(conn) {
	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	// output message received
	fmt.Print("Message Received:", string(message))
	// sample process for string received
	newmessage := strings.ToUpper(message)
	// send new string back to client
	conn.Write([]byte(newmessage + "\n"))
	conn.Close()
}
