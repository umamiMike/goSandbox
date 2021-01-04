package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

/*
marshal/unmarshl is about converting strings
encode/decode has to do with streams of data
*/

type Person struct {
	First string
	Last  string
	Age   int `json:"wisdom"`
}

func main() {
	p1 := Person{"Mike", "balssamic", 20}
	json.NewEncoder(os.Stdout).Encode(p1)

	var p2 Person                                                              // an itialized Person object
	rdr := strings.NewReader(`{"First":"James", "Last": "Bond", "wisdom":20}`) // this simulates a std input like I sent it from the command line or something
	json.NewDecoder(rdr).Decode(&p2)                                           //now decoding it like it was a stream input

	fmt.Println("-------")
	fmt.Println(p2.First) //proof by accessing values
	fmt.Println(p2.Last)
	fmt.Println(p2.Age) //shows the decoded values from the decoding

}
