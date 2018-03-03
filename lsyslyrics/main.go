package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var rules = map[string]string{
	"a": "ab",
	"b": "bac",
	"c": "ca",
}

type system struct {
	vars       []string
	constants  string
	axiom      string //the start, must be one of the keys in the rules
	iterations int
}

func (s system) run(r map[string]string) {
	// fmt.Println("vars", s.vars)
	// fmt.Println("constants", s.constants)
	// fmt.Println("axiom", s.axiom)
	// fmt.Println("iterations", s.iterations)
	teststring := s.axiom
	substring := ""
	for n := 1; n <= s.iterations; n++ {
		for _, r := range teststring {
			substring += rules[string(r)]
		}
		teststring += substring
		fmt.Println(" ")
		//	fmt.Println("teststring is: ", teststring)
	}
}

func printTimeCmd() *cobra.Command {
	return &cobra.Command{
		Use: "curtime",
		RunE: func(cmd *cobra.Command, args []string) error {
			now := time.Now()
			prettyTime := now.Format(time.RubyDate)
			cmd.Println("the current time is", prettyTime)
			return nil
		},
	}
}
func main() {
	cmd := &cobra.Command{
		Use:          "mike",
		Short:        "Hello",
		SilenceUsage: true,
	}
	cmd.AddCommand(printTimeCmd())
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
	sysrun := system{
		axiom:      "aba",
		iterations: 4,
	}
	sysrun.run(rules)
}
