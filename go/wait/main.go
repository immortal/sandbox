package main

import (
	"fmt"
	"os"
	"strconv"
	"syscall"
	"time"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Printf("usage: %s pid", os.Args[0])
		os.Exit(1)
	}

	pid, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		panic(err)
	}

	// On Unix systems, FindProcess always succeeds and returns a Process
	// for the given pid, regardless of whether the process exists.
	process, _ := os.FindProcess(int(pid))

	err = process.Signal(syscall.Signal(0))
	for err == nil {
		err = process.Signal(syscall.Signal(0))
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println(err)
}
