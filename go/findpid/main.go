package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"syscall"
)

func main() {
	for _, p := range os.Args[1:] {
		pid, err := strconv.ParseInt(p, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		process, err := os.FindProcess(int(pid))
		if err != nil {
			fmt.Printf("Failed to find process: %s\n", err)
		} else {
			err := process.Signal(syscall.Signal(0))
			if err != nil {
				fmt.Printf("process.Signal on pid %d returned: %v\n", pid, err)
			} else {
				fmt.Println("proccess exists")
			}
		}
	}
}
