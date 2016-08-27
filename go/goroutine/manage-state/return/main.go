package main

import (
	"fmt"
	"time"
)

var (
	ping = make(chan int)
	pong = make(chan int)
	quit = make(chan struct{})
	done = make(chan struct{})
)

func loopA() {
	for {
		select {
		case i := <-ping:
			fmt.Printf("loopA ping= %d\n", i)
		case <-quit:
			fmt.Println("after return will block")
			close(done)
			return
		}
	}
}

func loopB() {
	for i := 0; i <= 100000000000; i++ {
		ping <- i
	}
}

func exit() {
	for _ = range time.After(3 * time.Second) {
		close(quit)
	}
}

func main() {
	go loopA()
	go loopB()
	go exit()
	ping <- 0
	for _ = range done {
		fmt.Println("waiting")
	}
}
