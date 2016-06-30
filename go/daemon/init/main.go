package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"time"
)

var daemon = flag.Bool("daemonize", true, "-daemonize=true")

func init() {
	if !flag.Parsed() {
		flag.Parse()
	}
	if *daemon {
		args := os.Args[1:]
		args = append(args, "-daemonize=false")
		cmd := exec.Command(os.Args[0], args...)
		cmd.Start()
		fmt.Println("forking in PID: ", cmd.Process.Pid)
		os.Exit(0)
	}
}

func main() {

	parent := os.Getppid()
	child := os.Getpid()
	wd, _ := os.Getwd()

	pids := fmt.Sprintf("parent: %d, child: %d: cwd: %v", parent, child, wd)
	_ = ioutil.WriteFile("pids", []byte(pids), 0644)
	time.Sleep(100 * time.Second)
}
