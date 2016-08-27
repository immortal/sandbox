package main

import (
	"fmt"
	"time"
)

var (
	ping = make(chan int)
	quit = make(chan struct{})
)

func loopA() {
	for {
		select {
		case i := <-ping:
			fmt.Printf("loopA ping= %d\n", i)
			i++
		case <-quit:
			fmt.Println("after return will block")
			return
		}
	}
}

func loopB() {
	for _ = range time.After(3 * time.Second) {
		close(quit)
	}
}

func main() {
	go loopA()
	go loopB()
	for i := 0; i <= 10; i++ {
		ping <- i
		time.Sleep(1 * time.Second)
	}
}
