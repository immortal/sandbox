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

	go func() {
		cmd := exec.Command("sleep", "500")
		err = cmd.Start()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("sleep pid: %d", cmd.Process.Pid)
		err = cmd.Wait()
	}()

	log.Printf("Command finished with error: %v", err)

	waitForExit()
}
