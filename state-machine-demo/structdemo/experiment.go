// +build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/looplab/fsm"
)

type Work struct {
	To  string
	FSM *fsm.FSM
}

func newWork(to string) *Work {

	w := &Work{}
	w.FSM = fsm.NewFSM(
		"solid",
		fsm.Events{
			{Name: "melt", Src: []string{"solid"}, Dst: "liquid"},
			{Name: "freeze", Src: []string{"liquid"}, Dst: "solid"},
			{Name: "vaporize", Src: []string{"solid"}, Dst: "gas"},
			{Name: "condensce", Src: []string{"gas"}, Dst: "liquid"},
		},
		fsm.Callbacks{
			"enter_state": func(e *fsm.Event) { w.enterState(e) },
			"leave_state": func(e *fsm.Event) { w.leaveState(e) },
		},
	)
	return w
}

func (d *Work) leaveState(e *fsm.Event) {
	time.Sleep(1 * time.Millisecond)
	fmt.Println("you are about to leave:  ", e.Src)
}

func (d *Work) enterState(e *fsm.Event) {
	fmt.Printf("it is %s\n", e.Src)
	fmt.Println("you can ")
	for _, el := range d.FSM.AvailableTransitions() {
		fmt.Printf("%s,\n or ", el)
	}
}

func main() {
	work := newWork("poop")

	r := bufio.NewScanner(os.Stdin)

	fmt.Println("your current state is: ", work.FSM.Current())
	fmt.Println("you can ")
	for _, el := range work.FSM.AvailableTransitions() {
		fmt.Printf("%s,\n or ", el)
	}

	for {
		time.Sleep(1 * time.Second)

		for r.Scan() {
			text := r.Text()
			time.Sleep(1 * time.Millisecond)

			err := work.FSM.Event(text)
			if err != nil {
				fmt.Println("error:", err)
			}
			fmt.Print("What is your next move?:")
		}
	}
}
