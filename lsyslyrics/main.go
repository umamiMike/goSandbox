/**
I am working on a go program to generate L-systems
currently for strings only.

rules should be a slice of key value pairs
i need a function that recursively generates, and adds to the previous iteration

Vars: A,B,C,F,-,+ etc
One of the Vars acts as the axiom.  The trunk of the tree.

Example Rule
A: "

*/
package main

import (
	"fmt"
)

type system struct {
	Vars []string
}

func (s system) run() {
	for _, v := range s.Vars {
		fmt.Println("a variable is ", v)
	}
}

func main() {

	system_to_run := system{Vars: []string{"A", "B"}}
	fmt.Println(system_to_run)
	constants := "none"
	fmt.Println("constants are", constants)
	initiator := system_to_run.Vars[0]
	fmt.Println(initiator, "is the initiator")
	system_to_run.run()
}
