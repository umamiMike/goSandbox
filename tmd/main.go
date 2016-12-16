package main

//cli usage
// transposeMarkdown inputFile outputFile

import (
	"os"
	"strings"
	//	"fmt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	infile := os.Args[1]
	data := os.Args[2:]
	collapsedData := (strings.Join(data, " "))
	f, err := os.OpenFile(infile, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(collapsedData); err != nil {
		panic(err)
	}
	check(err)
}
