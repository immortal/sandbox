package main

import (
	"fmt"
	"os"
	"os/signal"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Printf("%s - %d Parent: %d Pid: %d\n", from, i, os.Getppid(), os.Getpid())
		if i > 3 {
			return
		}
	}
}

func waitForExit() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
	os.Exit(0)
}

func main() {

	pid := os.Getpid()

	parentpid := os.Getppid()

	fmt.Printf("Parent: %d pid: %d ", pid, parentpid)

	proc, err := os.FindProcess(pid)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Process: %d\n", proc.Pid)

	go f("goroutine")

	waitForExit()
}
