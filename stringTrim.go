package main

import "strings"
import "fmt"

func main() {
	fmt.Println("vim-go")
	lcn := "45675:55555"
	part := strings.Split(lcn, ":")[1]
	fmt.Println(part)

}
