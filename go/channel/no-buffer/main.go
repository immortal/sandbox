package main

import (
	"fmt"
	"time"
)

func spawn(i int, ch chan<- int) {
	ch <- i
}
func main() {
	ch := make(chan int)

	for i := 0; i <= 9; i++ {
		go spawn(i, ch)
	}

	for {
		select {
		case i := <-ch:
			if i == 7 {
				close(ch)
				return
			}
			fmt.Println(i)
			time.Sleep(time.Second)
			go spawn(i, ch)
		}
	}
}
