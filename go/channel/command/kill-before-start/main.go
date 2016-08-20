package main

import "os/exec"

func main() {
	cmd := exec.Command("sleep", "10")

	started := make(chan struct{}, 1)
	if err := cmd.Start(); err != nil {
		panic(err)
	}
	started <- struct{}{}

	go func() {
		<-started
		cmd.Process.Kill()
	}()

	cmd.Wait()
}
