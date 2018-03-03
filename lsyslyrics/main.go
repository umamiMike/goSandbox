package main

import (
	"fmt"
)

var rules = map[string]string{
	"A": "CASH_",
	"B": "BACK_",
	"C": "BibbleCob_",
}

type system struct {
	vars       []string
	constants  string
	axiom      string //the start, must be one of the keys in the rules
	iterations int
}

func (s system) run(r map[string]string) {
	fmt.Println("vars", s.vars)
	fmt.Println("constants", s.constants)
	fmt.Println("axiom", s.axiom)
	fmt.Println("iterations", s.iterations)
	teststring := s.axiom
	substring := ""
	for n := 1; n <= s.iterations; n++ {
		for _, r := range teststring {
			substring += rules[string(r)]
		}
		teststring += substring
	}
	fmt.Println("teststring is: ", teststring)
}

func main() {

	sysrun := system{
		vars:       []string{"A", "B"},
		axiom:      "A",
		iterations: 5,
	}
	sysrun.run(rules)
}
