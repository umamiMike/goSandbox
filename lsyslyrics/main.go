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
	Vars       []string
	Initiator  string
	Iterations int
}

//type ruleset struct {
//
//}
//func makeRulesets() *Ruleset {
//
//}

func (s system) run(r map[string]string) {
	for _, v := range s.Vars {
		fmt.Println("a variable is ", v)
		fmt.Println("and the value of ", r["A"])
	}
	output := s.Initiator

	for i := 1; i < s.Iterations; i++ {
		//		output = output +

		fmt.Println(output)
	}
}

func main() {

	systemliteral := system{
		Vars:       []string{"A", "B"},
		Initiator:  "A",
		Iterations: 3,
	}

	fmt.Println(systemliteral)
	constants := "none"
	fmt.Println("constants are", constants)
	initiator := systemliteral.Vars[0]
	ruleset := map[string]string{
		"A": "AB",
		"B": "BA",
	}
	fmt.Println(initiator, "is the initiator")
	fmt.Println(ruleset)

	systemliteral.run(ruleset)
}
