package main

import (
	"os/exec"
	"sync"
)

func main() {
	var lock sync.Mutex
	cmd := exec.Command("sleep", "10")

	lock.Lock()
	if err := cmd.Start(); err != nil {
		panic(err)
	}
	lock.Unlock()
	go func() {
		cmd.Process.Kill()
	}()

	cmd.Wait()
}
