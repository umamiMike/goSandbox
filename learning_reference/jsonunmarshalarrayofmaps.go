package main

import (
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
)

type Config struct {
	Dbs []map[string]string
}

func main() {
	encoded := `{ "dbs": [	
  { "test":  "string test"},
  { "new":  "string new"},
  { "dev1":  "string dev1"}
]}
`
	// Decode the json object
	u := &Config{}
	err := json.Unmarshal([]byte(encoded), &u)
	if err != nil {
		panic(err)
	}
	// Print out mother and father
	for i, num := range u.Dbs {
		fmt.Println(i, num)
		for k, v := range num {
			if k == "dev1" {
				fmt.Println("key:", k, "\n", v)
			}
		}
	}
}
