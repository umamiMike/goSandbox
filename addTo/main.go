package main

/*
cli usage
will add a string to a file of your choosing
I use it to add thoughts to various journals I keep.
addto the/file/path the string you want to add to the file.
*/
import (
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	infile := os.Args[1]
	if len(os.Args) > 1 {
		data := os.Args[2:] //everything after the first arg becomes data
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
}
