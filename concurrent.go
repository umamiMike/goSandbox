package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type result struct {
	URL string
	Dur time.Duration
}

func getter(url string, ch chan<- result) {
	defer close(ch)
	start := time.Now()
	if _, err := http.Get(url); err != nil {
		ch <- result{URL: url, Dur: time.Duration(-1)}
		return
	}
	ch <- result{URL: url, Dur: time.Now().Sub(start)}
}

func main() {
	urls := []string{"https://google.com", "https://yahoo.com", "https://bing.com", "https://duckduckgo.com"}
	durCh := make(chan result) //duration channel of type result
	var wg sync.WaitGroup      //create waitgroup
	wg.Add(len(urls))          //adds int which is the number of goroutines
	for _, url := range urls {
		_ = "breakpoint"
		ch := make(chan result)
		go getter(url, ch)
		go func() {
			defer wg.Done() //once the func is complutely done, it communicates that it is done to the waitgroup
			t := <-ch
			durCh <- t
		}()
	}
	go func() {
		wg.Wait()
		close(durCh)
	}()
	for t := range durCh {
		_ = "breakpoint"
		fmt.Println(t)
	}
}
