package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func waitForExit() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
	os.Exit(0)
}

func main() {

	for i := 1; i < 1000; i++ {
		if i%3 == 0 {
			fmt.Fprintf(os.Stderr, "STDERR i: %d\n", i)
		} else {
			fmt.Printf("STDOUT i: %d\n", i)
		}
		time.Sleep(1 * time.Second)
	}

}
