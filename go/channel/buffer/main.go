package main

import (
	//	"fmt"
	"os/exec"
)

func main() {
	done := make(chan error, 1) // buffer because goroutine might write to done<- before main goroutine starts to wait on <-done
	//	done <- fmt.Errorf("2")
	cmd := exec.Command("ls")
	go func() {
		err := cmd.Wait()
		select {
		case done <- err:
		default:
		}
	}()
	// don't panic if buffered
	//	err := <-done
	//	fmt.Println(err)
}
