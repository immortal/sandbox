// to test the panic use go build -race
package main

import "fmt"

func main() {
	go run()
	for i := 0; i <= 10; i++ {
		signal <- struct{}{}
		if <-read {
			ping <- "ping"
		}
		if i%3 == 0 {
			write <- true
		}
	}
}

func run() {
	foo := false
	for {
		select {
		case <-signal:
			fmt.Println("signal", foo)
			read <- foo
		case foo = <-write:
			fmt.Println("write", foo)
		case v := <-ping:
			fmt.Println(v, foo)
			foo = false
		}
	}
}

var (
	ping   = make(chan string)
	signal = make(chan struct{})
	read   = make(chan bool)
	write  = make(chan bool)
)
