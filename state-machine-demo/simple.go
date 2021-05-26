package main

import (
	"fmt"

	"github.com/looplab/fsm"
)

type Door struct {
	To  string
	FSM *fsm.FSM
}

func NewDoor(to string) *Door {
	d := &Door{
		To: to,
	}

	d.FSM = fsm.NewFSM(
		"closed", //initial state
		fsm.Events{
			{Name: "open", Src: []string{"closed"}, Dst: "open"},
			{Name: "close", Src: []string{"open"}, Dst: "closed"},
			{Name: "close", Src: []string{"open"}, Dst: "closed"},
		},
		fsm.Callbacks{
			"enter_state": func(e *fsm.Event) { d.enterState(e) },
			"leave_state": func(e *fsm.Event) { d.leaveState(e) },
			"before_open": func(e *fsm.Event) { d.beforeOpen(e) },
		},
	)

	return d
}

func (d *Door) enterState(e *fsm.Event) {
	fmt.Printf("The door to %s is %s\n", d.To, e.Dst)
}

func (d *Door) leaveState(e *fsm.Event) {
	fmt.Printf("leaving %s and on to %s\n", e.Src, e.Dst)
}

func (d *Door) beforeOpen(e *fsm.Event) {
	fmt.Printf("the door to %s is %s, and will soon %s\n\n", d.To, e.Src, e.Dst)
}

func main() {
	door := NewDoor("hell")

	err := door.FSM.Event("open")
	if err != nil {
		fmt.Println(err)
	}

	err = door.FSM.Event("close")
	if err != nil {
		fmt.Println(err)
	}
}
