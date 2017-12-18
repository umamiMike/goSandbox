package main

import (
	"fmt"
	//"gopkg.in/gomail.v2"
	"os"
	"os/exec"
	"strings"
)

func main() {
	exec.Command("source ./conf.env")
	user := os.Getenv("E_USER")
	//pass := os.Getenv("E_PASS")
	fmt.Println(user)
}
