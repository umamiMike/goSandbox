package main

//cli usage
// transposeMarkdown inputFile outputFile


import (
	"os"
	"io/ioutil"
	//"bufio"
	//"fmt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}




func main(){

	infile := os.Args[1]
	data := []byte(os.Args[2])

	err := ioutil.WriteFile(infile, data,0644)
	check(err)
}