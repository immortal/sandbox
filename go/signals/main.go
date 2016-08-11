package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	fmt.Printf("Pid = %+v\n", os.Getpid())
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	select {
	// kill -2 , kill -INT
	case s := <-c:
		fmt.Printf("Got signal: %d", s)
		os.Exit(0)
	case <-time.After(30 * time.Second):
		os.Exit(1)
	}
}
