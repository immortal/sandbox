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

	if err := exec.CommandContext(ctx, "sleep", "5").Run(); err != nil {
		// will be interrupted.
	}
}
