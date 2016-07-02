package main

import (
	"fmt"
	"os"
	"strconv"
	"syscall"
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

	err = syscall.PtraceAttach(process.Pid)
	if err != nil {
		fmt.Printf("err: %s", err)
	}
}
