// +build ignore

package main

import (
	"bufio"
	"fmt"
	"github.com/looplab/fsm"
	"os"
	"time"
)

/*
 *type Event struct {
 *    // FSM is a reference to the current FSM.
 *    FSM *FSM
 *
 *    // Event is the event name.
 *    Event string
 *
 *    // Src is the state before the transition.
 *    Src string
 *
 *    // Dst is the state after the transition.
 *    Dst string
 *
 *    // Err is an optional error that can be returned from a callback.
 *    Err error
 *
 *    // Args is a optinal list of arguments passed to the callback.
 *    Args []interface{}
 *    // contains filtered or unexported fields
 *}
 *
 */
type Work struct {
	To  string
	FSM *fsm.FSM
}

func newWork(to string) *Work {

	w := &Work{}
	w.FSM = fsm.NewFSM(
		"idle",
		fsm.Events{
			{Name: "start", Src: []string{"stop"}, Dst: "start"},
			{Name: "stop", Src: []string{"start", "stop", "continue"}, Dst: "stopped"},
			{Name: "continue", Src: []string{"start", "stopped", "continue"}, Dst: "stop"},
		},
		fsm.Callbacks{
			"enter_state": func(e *fsm.Event) { w.enterState(e) },
			"leave_state": func(e *fsm.Event) { w.leaveState(e) },
			"leave_stop": func(e *fsm.Event) {
				w.leaveStop(e)
			},
		},
	)
	return w
}
func (d *Work) leaveStop(e *fsm.Event) {

	fmt.Println(`
*************************
you totally entered the stop state
*************************
`)
}

func (d *Work) leaveState(e *fsm.Event) {
	fmt.Println("you are about to leave:  ", e.Src)
	fmt.Println("you will be in ", e.Event)
}

func (d *Work) enterState(e *fsm.Event) {
	fmt.Println("you were in ", e.Src)
	fmt.Println("now you are in ", e.Event)
	fmt.Printf("and the destination is:   %s\n", e.Dst)
}

func main() {
	work := newWork("poop")

	r := bufio.NewScanner(os.Stdin)

	fmt.Println("your current state is: ", work.FSM.Current())

	for {
		time.Sleep(1 * time.Second)

		for r.Scan() {
			text := r.Text()
			time.Sleep(1 * time.Millisecond)

			err := work.FSM.Event(text)
			if err != nil {
				fmt.Println("error:", err)
			}
			fmt.Println("\n\nyour current state is: ", work.FSM.Current())
			fmt.Println("available transitios are: ", work.FSM.AvailableTransitions())
			fmt.Print("What is your next move?:")

		}

	}

}
