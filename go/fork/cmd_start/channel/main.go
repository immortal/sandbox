package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
)

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

	cmd := exec.Command("sleep", "500")

	started := make(chan error)
	done := make(chan error)
	go func() {
		started <- cmd.Start()
		done <- cmd.Wait()
	}()

	<-started
	println("Started")
	log.Printf("sleep pid: %d", cmd.Process.Pid)

	err = <-done
	println("Done")
	if err != nil {
		fmt.Printf("wait err: %s\n", err.Error())
	}
	waitForExit()
}
