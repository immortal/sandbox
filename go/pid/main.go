package main

import (
	"fmt"
	"os"
)

func main() {

	pid := os.Getpid()

	parentpid := os.Getppid()

	fmt.Printf("The parent process id of %v is %v\n", pid, parentpid)

	proc, err := os.FindProcess(pid) // or replace with other process number

	if err != nil {
		panic(err)
	}

	fmt.Println(proc.Pid)

}
