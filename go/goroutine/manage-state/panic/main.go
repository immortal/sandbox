// to test the panic use go build -race
package main

import "fmt"

type test struct {
	ch  chan string
	foo bool
}

func (t *test) run() {
	for {
		select {
		case v := <-t.ch:
			fmt.Printf("%+v, foo=%+v\n", v, t.foo)
			t.foo = false
		default:
		}
	}
}

func (self *test) Ping() {
	self.ch <- "ping"
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
		if t.foo {
			t.Ping()
		}
		//	time.Sleep(time.Second)
		if i%3 == 0 {
			t.foo = true
		}
	}
}
