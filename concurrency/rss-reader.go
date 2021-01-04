package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

type Fetcher interface { 
Fetch() (items []Item, next time.Time, err error)
}

func Fetch(doman string) Fetcher {...} //fetches Items from domain

type Subscription interface {
	Updates() <-chan Item //stream of Items
	Close() error //shuts down the stream
}
func Subscribe(fetcher Fetcher) Subscription {...} //converts Fetches to stream
func Merge(subs ...Subscription) Subscription {...} // merges several streams

