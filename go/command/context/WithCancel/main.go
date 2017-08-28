package main

import (
	"context"
	"os/exec"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func(c context.CancelFunc) {
		time.Sleep(time.Second)
		c()
	}(cancel)

	cmd := exec.CommandContext(ctx, "sleep", "5")
	cmd.Start()
	cmd.Wait()
}
