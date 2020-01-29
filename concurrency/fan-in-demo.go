package main

import (
	"fmt"
	"github.com/icrowley/fake"
	"math/rand"
	"time"
)

func producer(ch chan string, t1 time.Time, name string) {
	var i int
	r := rand.New(rand.NewSource(99))
	for {
		t2 := time.Now()
		d := t2.Sub(t1)
		s := d.Seconds()
		ch <- name + "," + fmt.Sprintf("%v", s)
		i++
		rantime := time.Duration(r.Intn(1)) * time.Second
		time.Sleep(rantime)
	}
}

func reader(out chan string) {
	for x := range out {
		fmt.Println(x)
	}
}

func main() {

	ch := make(chan string)
	out := make(chan string)
	fmt.Println("name, seconds")
	for i := 0; i < 200; i++ {

		go producer(ch, time.Now(), fake.FirstName())
	}
	go reader(out)
	for i := range ch {
		out <- i
	}
}
