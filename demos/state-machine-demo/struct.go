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
		"closed",
		fsm.Events{
			{Name: "open", Src: []string{"closed"}, Dst: "opened"}, //use the taxonomy of the Dst: in the Src strings
			{Name: "close", Src: []string{"opened", "splayed"}, Dst: "closed"},
			{Name: "splay", Src: []string{"closed"}, Dst: "splayed"},
		},
		fsm.Callbacks{
			"enter_state": func(e *fsm.Event) { d.enterState(e) },
		},
	)

	return d
}

func (d *Door) enterState(e *fsm.Event) {
	fmt.Printf("The door to %s is %s\n", d.To, e.Dst)
}

func main() {
	door := NewDoor("heaven")

	fmt.Println("The current state of the door is: ", door.FSM.Current(), ", can the door open?", door.FSM.Can("open"))
	err := door.FSM.Event("open")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("The current state of the door is: ", door.FSM.Current(), ", can the door close?", door.FSM.Can("close"))
	err = door.FSM.Event("close")
	if err != nil {
		fmt.Println(err)
	}
	err = door.FSM.Event("splay")
	if err != nil {
		fmt.Println(err)
	}
}
