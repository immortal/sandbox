package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	// replace  io.Pipe() with os.Pipe() to make it work
	//	pr, pw, := io.Pipe()
	pr, pw, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("/bin/bash", "-c", "cat")
	cmd.Stdin = pr

	started := make(chan error)
	done := make(chan error)
	go func() {
		started <- cmd.Start()
		done <- cmd.Wait()
	}()

	<-started
	println("Started")
	pw.Write([]byte("hello"))

	cmd.Process.Signal(syscall.SIGTERM)

	err = <-done
	println("Done")
	if err != nil {
		fmt.Printf("wait err: %s\n", err.Error())
	}
}
