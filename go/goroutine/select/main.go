package main

import (
	"fmt"
	"os/exec"
	"time"
)

type Exec struct {
	ch chan error
}

func (e *Exec) Proc() *Exec {
	cmd := exec.Command("sleep", "35")
	if err := cmd.Start(); err != nil {
		panic(err)
	}
	fmt.Printf("cmd.Process.Pid = %+v\n", cmd.Process.Pid)
	go func() {
		err := cmd.Wait()
		fmt.Printf("err = %+v\n", err)
		e.ch <- err
	}()
	return e
}

func main() {
	e := &Exec{}
	e.Proc()
	for {
		select {
		case err := <-e.ch:
			if exitError, ok := err.(*exec.ExitError); ok {
				fmt.Println(exitError)
			} else {
				fmt.Println("naranjas")
			}

		default:
			fmt.Println("default")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
