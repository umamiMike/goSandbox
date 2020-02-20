package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	p := make([]byte, 2048)
	conn, err := net.Dial("udp", "127.0.0.1:1234")
	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("input: ")
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Fprintf(conn, text)
		fmt.Print("input: ")

	}

	if err == nil {
		fmt.Printf("%s\n", p)
	} else {
		fmt.Printf("Some error %v\n", err)
	}
	conn.Close()
}
