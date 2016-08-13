package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("/tmp/stdout")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		panic(err)
	}

	if err := cmd.Start(); err != nil {
		panic(err)
	}

	println(cmd.Process.Pid)

	defer cmd.Wait()

	// Non-blockingly echo command output to terminal
	go io.Copy(os.Stdout, stdout)
	go io.Copy(os.Stderr, stderr)

	fmt.Println("...")
}
