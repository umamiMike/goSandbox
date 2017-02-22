package main

import (
	"fmt"
	"strings"
)

func main() {
	longString := "I am a long string"
	fmt.Println(strings.IndexAny("long", longString))
}
