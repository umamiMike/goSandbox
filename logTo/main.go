package main

/*
cli usage
will add a string to a file of your choosing
I use it to add thoughts to various journals I keep.
addto the/file/path the string you want to add to the file.
TODO: add cobra commands for different props
create a .logto in home with list of files logged to

*/
import (
	"os"
	"strings"
	"time"
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
		if _, err = f.WriteString("\n\n" + time.Now().Format("01-02-2006 15:04:05") + " - " + collapsedData); err != nil {
			panic(err)
		}
		check(err)
	}
}
