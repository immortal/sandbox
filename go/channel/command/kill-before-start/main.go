package main

import "os/exec"

func main() {
	cmd := exec.Command("sleep", "10")
	started := make(chan struct{}, 1)

	go func(cmd *exec.Cmd, signal chan struct{}) {
		if err := cmd.Start(); err != nil {
			panic(err)
		}
		started <- struct{}{}
	}(cmd, started)

	<-started
	cmd.Process.Kill()
}
