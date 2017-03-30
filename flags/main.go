package main

import (
	"flag"
	"fmt"
)

func main() {
	hostPtr := flag.String("host", "localhost", "a string of the host")
	flag.Parse() //run to process he the flags into their variables
	fmt.Println("host is", *hostPtr)

}
