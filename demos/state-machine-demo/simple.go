package main

import (
	"fmt"
	"github.com/looplab/fsm"
)

func main() {
	fsm := fsm.NewFSM(
		"closed",
		fsm.Events{
			{Name: "open", Src: []string{"closed"}, Dst: "open"},
			{Name: "splay", Src: []string{"closed"}, Dst: "splayed"},
			{Name: "close", Src: []string{"open", "splay"}, Dst: "closed"},
		},
		fsm.Callbacks{},
	)
	fmt.Println(fsm.Current())
	err := fsm.Event("open")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()
	fmt.Println("The current state of the door is: ", fsm.Current(), ", can the door close?", fsm.Can("close"))
	err = fsm.Event("close")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("The current state of the door is: ", fsm.Current(), ", can the door close?", fsm.Can("close"))

}
