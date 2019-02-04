// +build ignore

package main

import (
	"bufio"
	"fmt"
	"github.com/looplab/fsm"
	"os"
	"time"
)

func main() {
	fsm := fsm.NewFSM(
		"idle",
		fsm.Events{
			{Name: "scan", Src: []string{"idle"}, Dst: "scanning"},
			{Name: "working", Src: []string{"scanning"}, Dst: "scanning"},
			{Name: "situation", Src: []string{"scanning"}, Dst: "scanning"},
			{Name: "situation", Src: []string{"idle"}, Dst: "idle"},
			{Name: "finish", Src: []string{"scanning"}, Dst: "idle"},
		},
		fsm.Callbacks{
			"scan": func(e *fsm.Event) {
				fmt.Println("after_scan: " + e.FSM.Current())
			},
			"working": func(e *fsm.Event) {
				fmt.Println("working: " + e.FSM.Current())
			},
			"situation": func(e *fsm.Event) {
				fmt.Println("situation: " + e.FSM.Current())
			},
			"finish": func(e *fsm.Event) {
				fmt.Println("finish: " + e.FSM.Current())
			},
		},
	)
	r := bufio.NewScanner(os.Stdin)
	fmt.Println(fsm.Current())
	for {
		time.Sleep(1 * time.Second)
		err := fsm.Event("scan")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print("what is the next event? ")

		for r.Scan() {
			text := r.Text()
			time.Sleep(3 * time.Second)
			fmt.Println("3:" + fsm.Current())

			err = fsm.Event(text)
			if err != nil {
				fmt.Println(err)
			}

		}

	}

}
