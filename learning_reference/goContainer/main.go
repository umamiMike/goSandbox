/**
https://youtu.be/HPuvDm8IC-4
*/
package main

import (
	"fmt"
	"os"
	"os/exec"
	//	"syscall"
)

// docker run <container> command args
//go run main.go run command args
func main() {
	switch os.Args[1] {
	case "run":
		run()
	default:
		panic("what?")

	}

}

func run() {
	fmt.Println("vim-go")
	fmt.Printf("running %v\n", os.Args[2:])
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	//	cmd.SysProcAttr = &syscall.SysProcAttr{
	//		Cloneflags: syscall.CLONE_NEWUTS,
	//	}
	must(cmd.Run())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
