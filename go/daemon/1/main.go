package main

import (
	"flag"
	"fmt"
	"os"
	"syscall"
	"time"
)

var isChild = flag.Bool("child", false, "parent pid")

func main() {
	flag.Parse()
	var s string
	if !*isChild { //parent
		fmt.Println("forking")
		files := make([]*os.File, 3)
		files[syscall.Stdin] = os.Stdin
		files[syscall.Stdout] = os.Stdout
		files[syscall.Stderr] = os.Stderr
		fmt.Println(os.StartProcess(os.Args[0], append(os.Args, "-child"), &os.ProcAttr{
			Dir:   "/tmp",
			Env:   os.Environ(),
			Files: files,
		}))
		fmt.Scan(&s) // block
	} else {
		ppid := os.Getppid()
		fmt.Println("ppid", ppid, "kill:", syscall.Kill(ppid, syscall.SIGTERM))
		time.Sleep(5 * time.Second)
		fmt.Println("child dying", ppid)
	}
}
