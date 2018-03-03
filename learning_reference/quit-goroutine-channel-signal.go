package main

import "sync"
import "time"
import "math/rand"

var wg sync.WaitGroup

func main() {
	wg.Add(1)

	ch := make(chan int)
	go SomeGoRoutine(ch) //feed it the channel when running

	ch <- 1
	ch <- 8
	ch <- 2
	ch <- 3
	close(ch)

	wg.Wait()
}
func SomeGoRoutine(ch chan int) {
	for {
		foo, ok := <-ch //reads in from the channel
		r := rand.Intn(100)
		println("r is: ", r)
		time.Sleep(time.Duration(r) * time.Millisecond)
		if foo == 8 {
			println("you caught the snitch eight ball")
		}

		if !ok {
			println("done")
			wg.Done()
			return
		}
		println(foo)
	}
}
