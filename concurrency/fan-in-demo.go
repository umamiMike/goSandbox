package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

/*
 * takes a channel and a duration
 inits an interger,
 writes to channel
 increments integer
 sleeps for a random amount of time in seconds
*/
func producer(ch chan string, d time.Duration, name string) {
	var i int
	r := rand.New(rand.NewSource(99))
	for {
		ch <- strconv.Itoa(i) + name
		i++
		rantime := time.Duration(r.Intn(2)) * time.Second
		time.Sleep(rantime + d)
	}
}

/*
 * takes an int channel and prints out x from the channel indefinitely

 */
func reader(out chan string) {
	for x := range out {
		fmt.Println(x)
	}
}

/*
 * makes to channels
 * spawns 3 producers
 */
func main() {

	ch := make(chan string)
	out := make(chan string)
	go producer(ch, 10*time.Millisecond, "barf")
	go producer(ch, 25*time.Millisecond, "bag")
	go reader(out)
	for i := range ch {
		out <- i
	}
}
