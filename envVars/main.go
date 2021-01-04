package main

import "fmt"
import "os"
import "syscall"
import "strings"

func main() {
	fmt.Println("vim-go")
	os.Setenv("FOO", "1")
	fmt.Println("FOO", os.Getenv("FOO"))
	fmt.Println("BAR", os.Getenv("BAR"))
	os.Unsetenv("BAR")
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair)
	}
	syscall.Exec(os.Getenv("SHELL"), []string{os.Getenv("SHELL")}, syscall.Environ())
}
