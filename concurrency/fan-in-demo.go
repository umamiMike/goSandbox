package main

import (
	"fmt"
	"math/rand"
	"time"
)

func producer(ch chan int, d time.Duration) {
	var i int
	r := rand.New(rand.NewSource(99))
	for {
		ch <- i
		i++
		rantime := time.Duration(r.Intn(6)) * time.Second
		fmt.Println("I sleep for this long:", rantime)
		time.Sleep(rantime + d)
	}
}

func reader(out chan int) {
	for x := range out {
		fmt.Println(x)
	}
}

func main() {

	ch := make(chan int)
	out := make(chan int)
	go producer(ch, 100*time.Millisecond)
	go producer(ch, 250*time.Millisecond)
	go reader(out)
	for i := range ch {
		out <- i
	}
}
