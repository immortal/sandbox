package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

func main() {
	file, err := os.Create("Lock")
	if err != nil {
		fmt.Printf("Create: %s\n", err)
		return
	}
	err = syscall.Flock(int(file.Fd()), syscall.LOCK_EX)
	if err != nil {
		file.Close()
		fmt.Printf("Flock: %s\n", err)
		return
	}
	fmt.Printf("pid: %d  sleeping...", os.Getpid())
	time.Sleep(60 * time.Second)
}
