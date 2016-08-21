package main

import (
	"fmt"
	"time"
)

type MyType struct{}
type MyOtherType struct{}

type test struct {
	ch chan interface{}
}

func (self *test) run() {

	for ch := range self.ch {
		switch ch := ch.(type) {
		case MyType:
			fmt.Printf("%#v\n", ch)
		case MyOtherType:
			fmt.Printf("%#v\n", ch)
		}
	}
}

func (self *test) TestMyType() {
	self.ch <- MyType{}
}
func (self *test) TestMyOtherType() {
	self.ch <- MyOtherType{}
}

func New() *test {
	t := &test{
		ch: make(chan interface{}),
	}
	go t.run()
	return t
}

func main() {
	t := New()
	for i := 0; i <= 10; i++ {
		if i%3 == 0 {
			t.TestMyType()
		} else {
			t.TestMyOtherType()
		}
		time.Sleep(time.Second)
	}
}
