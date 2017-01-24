package main

import (
	"encoding/json"
	"fmt"
)

var jsonStr string = `{"title":"Buy cheese and bread for breakfast."}`

func main() {
	var parsed map[string]interface{}
	json.Unmarshal([]byte(jsonStr), &parsed)
	fmt.Println(parsed)
}
