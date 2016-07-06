package main

import (
	"fmt"
	"time"
)

func sleep(ch chan<- bool) {
	go func() {
		time.Sleep(1 * time.Second)
		ch <- true
	}()
}
func quit(ch chan<- struct{}) {
	go func() {
		time.Sleep(3 * time.Second)
		close(ch)
	}()
}

func main() {

	timeout := make(chan bool)
	q := make(chan struct{})

	sleep(timeout)
	quit(q)

	for {
		select {
		case <-q:
			return
		case <-timeout:
			fmt.Println("timeout true")
			sleep(timeout)
		}
	}
}
