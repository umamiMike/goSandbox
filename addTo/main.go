package main

//cli usage
// addto the/file/path the string you want to add to the file.

import (
	"fmt"
	"os"
	"strings"
)

func myPrinter(i) {
	fmt.Printf("%T", i)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	infile := os.Args[1]
	data := os.Args[2:] //everything after the first arg becomes data

	// 	data2 := make(string "\n", data)

	collapsedData := (strings.Join(data, " "))
	f, err := os.OpenFile(infile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(collapsedData); err != nil {
		panic(err)
	}
	check(err)
}
