package main

import (
	"context"
	"log"
	"os/exec"
	"time"
)

func Run(quit chan struct{}) {
	ctx, cancel := context.WithCancel(context.Background())
	cmd := exec.CommandContext(ctx, "sleep", "300")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		log.Println("waiting cmd to exit")
		err := cmd.Wait()
		if err != nil {
			log.Println(err)
		}
	}()

	go func() {
		select {
		case <-quit:
			log.Println("calling ctx cancel")
			cancel()
		}
	}()
}

func main() {
	ch := make(chan struct{})
	Run(ch)
	select {
	case <-time.After(3 * time.Second):
		log.Println("closing via ctx")
		ch <- struct{}{}
	}
}
