package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"syscall"
	"time"
)

func init() {
	if os.Getppid() > 1 {
		cmd := exec.Command(os.Args[0])
		cmd.Start()
		fmt.Println("forking in PID: ", cmd.Process.Pid)
		os.Exit(0)
	}

	os.Chdir("/")
	_ = syscall.Umask(0)
	//	create a new SID for the child process
	_, s_errno := syscall.Setsid()
	if s_errno != nil {
		log.Printf("Error: syscall.Setsid errno: %d", s_errno)
	}
	fmt.Println("Parent PID: ", os.Getppid())
}

func main() {

	parent := os.Getppid()
	child := os.Getpid()
	wd, _ := os.Getwd()

	pids := fmt.Sprintf("parent: %d, child: %d: cwd: %v", parent, child, wd)
	_ = ioutil.WriteFile("/tmp/pids", []byte(pids), 0644)
	time.Sleep(60 * time.Second)
}
