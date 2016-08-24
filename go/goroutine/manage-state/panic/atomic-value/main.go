// to test the panic use go build -race
package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type test struct {
	ch chan string
	atomic.Value
}

func (t *test) run() {
	for {
		select {
		case v := <-t.ch:
			fmt.Printf("%+v, foo=%+v\n", v, t.Load())
			t.Store(false)
		case <-time.After(time.Second):
			fmt.Println("running task")
		}
	}
}

func (t *test) Ping() {
	t.ch <- "ping"
}

func New() *test {
	t := &test{
		ch: make(chan string),
	}
	go t.run()
	return t
}

func main() {
	t := New()
	for i := 0; i <= 10; i++ {
		if x, _ := t.Load().(bool); x {
			t.Ping()
		}
		//	time.Sleep(time.Second)
		if i%3 == 0 {
			t.Store(true)
		}
	}
}
