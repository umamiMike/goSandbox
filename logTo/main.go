package main

/*
cli usage
will add a string to a file of your choosing
I use it to add thoughts to various journals I keep.
addto the/file/path the string you want to add to the file.
TODO: add cobra commands for different props
TODO: create a .logto in home with list of files logged to
TODO: find backward from the end to see if the date portion of the log exists and is the same
	if yes
		only add the time stamp
	else do what already does
TODO: make time part prettier format


*/
import (
	"fmt"
	"os"
	"strings"
	"text/template"
	"time"
)

const outputTemplate = `
{{ .Prefix  }} {{.Time }}
{{.Output }}`

var infile = os.Args[1]
var data = os.Args[2:]

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	if len(data) == 0 { //no message sent, only the file...nothing to write
		os.Exit(3)
	}
	collapsedData := (strings.Join(data, " "))
	f, err := os.OpenFile(infile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	pre := "#"
	write(f, pre, collapsedData)
	defer f.Close()
}
func write(f *os.File, pre string, data string) {

	output := struct {
		Prefix string
		Time   string
		Output interface{}
	}{
		pre,
		time.Now().Format("01-02-2006 15:04:05"),
		data,
	}
	t, err := template.New("output").Parse(outputTemplate)
	if err != nil {
		panic(err)
	}
	t.Execute(f, output)
	if err != nil {
		panic(err)
	}

}
