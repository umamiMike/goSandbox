package main

import (
	"fmt"
	"github.com/satori/go.uuid"
)

func main() {
	// Creating UUID Version 4
	u1 := uuid.NewV4()
	fmt.Printf("UUIDv4: %d\n", u1)

	// Parsing UUID from string input
	u2, err := uuid.FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	if err != nil {
		fmt.Printf("Something gone wrong: %s", err)
	}
	fmt.Printf("Successfully parsed: %s", u2)

	//	u3 := uuid.NewV5()
	//	fmt.Printf("UUIDv5: %s\n", u3)

	// mysql likes binary(16) encodeing so probably need to do something like
	//func (u UUID) MarshalBinary() (data []byte, err error)
}
