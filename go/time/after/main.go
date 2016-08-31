package main

import (
	"fmt"
	"os/exec"
	"time"
)

func Exec(done chan<- error) error {
	cmd := exec.Command("./start")
	if err := cmd.Start(); err != nil {
		return err
	}
	go func() {
		done <- cmd.Wait()
	}()
	return nil
}

func main() {
	var (
		run  = make(chan struct{}, 1)
		done = make(chan error, 1)
	)

	Exec(done)

	for {
		select {
		case <-run:
			err := Exec(done)
			if err != nil {
				fmt.Println(err)
				//	time.AfterFunc(3*time.Second, func() {
				time.Sleep(time.Second)
				run <- struct{}{}
				//	})
			}
		default:
			select {
			case err := <-done:
				fmt.Println(err)
				run <- struct{}{}
			}
		}
	}
}
