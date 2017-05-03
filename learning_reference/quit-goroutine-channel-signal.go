package main

import "sync"

var wg sync.WaitGroup

func main() {
	wg.Add(1)

	ch := make(chan int)
	go SomeGoRoutine(ch) //feed it the channel when running

	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	wg.Wait()
}
func SomeGoRoutine(ch chan int) {
	for {
		foo, ok := <-ch
		if !ok {
			println("done")
			wg.Done()
			return
		}
		println(foo)
	}
}
