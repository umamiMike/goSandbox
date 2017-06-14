package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var Ball int
	table := make(chan int)
	go player(table, "joe")
	go player(table, "sammantha")
	go player(table, "jimmy")

	table <- Ball
	time.Sleep(6 * time.Second)
	<-table
}

func player(table chan int, name string) {
	for {
		ball := <-table
		ball++
		fmt.Println(name)
		time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)
		table <- ball
	}
}
