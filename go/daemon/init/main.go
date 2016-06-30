package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"syscall"
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

	_ = syscall.Umask(0)
	os.Chdir("/")

	//	create a new SID for the child process
	_, s_errno := syscall.Setsid()
	if s_errno != nil {
		log.Printf("Error: syscall.Setsid errno: %d", s_errno)
	}

	os.Chdir("/")

}

func main() {

	parent := os.Getppid()
	child := os.Getpid()
	wd, _ := os.Getwd()

	pids := fmt.Sprintf("parent: %d, child: %d: cwd: %v", parent, child, wd)
	_ = ioutil.WriteFile("/tmp/pids", []byte(pids), 0644)
	time.Sleep(100 * time.Second)
}
