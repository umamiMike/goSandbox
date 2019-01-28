package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(tasksCh <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		task, ok := <-tasksCh
		if !ok {
			return
		}
		r := rand.New(rand.NewSource(99))
		rantime := time.Duration(r.Intn(2)) * time.Second
		time.Sleep(rantime)
		fmt.Println("processing task", task)
	}
}

func pool(wg *sync.WaitGroup, workers, tasks int) {

	tasksCh := make(chan int)

	for i := 0; i < workers; i++ {
		go worker(tasksCh, wg)
	}

	for i := 0; i < tasks; i++ {
		tasksCh <- i
	}

	close(tasksCh)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	go pool(&wg, 5, 50)
	wg.Wait()
}
