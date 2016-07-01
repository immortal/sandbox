package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"syscall"
	"time"
)

var daemon = flag.Int("daemonize", 1, "-daemonize=1")

func init() {
	if !flag.Parsed() {
		flag.Parse()
	}
	if *daemon > 0 {
		args := os.Args[1:]
		if *daemon == 1 {
			args = append(args, "-daemonize=2")
		} else if *daemon == 2 {
			i := 0
			for ; i < len(args); i++ {
				if args[i] == "-daemonize=2" {
					args[i] = "-daemonize=3"
					break
				}
			}
			//os.Chdir("/")
			_ = syscall.Umask(0)
			_, err := syscall.Setsid()
			if err != nil {
				panic(err)
			}
		} else if *daemon == 3 {
			i := 0
			for ; i < len(args); i++ {
				if args[i] == "-daemonize=3" {
					args[i] = "-daemonize=4"
					break
				}
			}
		} else if *daemon == 4 {
			i := 0
			for ; i < len(args); i++ {
				if args[i] == "-daemonize=4" {
					args[i] = "-daemonize=5"
					break
				}
			}
			//os.Chdir("/")
			_ = syscall.Umask(0)
			_, err := syscall.Setsid()
			if err != nil {
				panic(err)
			}
		} else if *daemon == 5 {
			i := 0
			for ; i < len(args); i++ {
				if args[i] == "-daemonize=5" {
					args[i] = "-daemonize=0"
					break
				}
			}
		}
		data := fmt.Sprintf("%v %v %v %#v", *daemon, os.Getpid(), os.Getppid(), args)
		_ = ioutil.WriteFile("/tmp/pids-"+fmt.Sprintf("%d", *daemon), []byte(data), 0644)
		cmd := exec.Command(os.Args[0], args...)
		cmd.Start()
		fmt.Println("forking in PID: ", cmd.Process.Pid, args)
		if *daemon == 1 {
			os.Exit(0)
		}
	}

	os.Chdir("/")
	data := fmt.Sprintf("%v %v %v", *daemon, os.Getpid(), os.Getppid())
	_ = ioutil.WriteFile("/tmp/pids-fin", []byte(data), 0644)
}

func main() {

	parent := os.Getppid()
	child := os.Getpid()
	wd, _ := os.Getwd()

	pids := fmt.Sprintf("parent: %d, child: %d: cwd: %v", parent, child, wd)
	_ = ioutil.WriteFile("/tmp/pids", []byte(pids), 0644)
	time.Sleep(300 * time.Second)
}
