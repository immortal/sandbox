package main

import (
	"fmt"
	"os/exec"
)

func main() {
	done := make(chan error, 1) // buffer because goroutine might write to done<- before main goroutine starts to wait on <-done
	cmd := exec.Command("ls")
	go func() {
		select {
		case done <- cmd.Wait():
		default:
		}
	}()
	cmd.Start()
	fmt.Printf("done = %+v\n", <-done)
}
