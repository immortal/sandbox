package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"
)

func main() {
	exit := make(chan error, 1)
	go run(exit)

	for {
		select {
		case <-exit:
			println("fin, restarting")
			run(exit)
		default:
			time.Sleep(time.Second)
			println("running...")
		}
	}
}

func run(ch chan<- error) {
	cmd := exec.Command("sleep", "3")
	//cmd.SysProcAttr = &syscall.SysProcAttr{
	//Setpgid: true,
	//}
	go func() {
		if err := cmd.Start(); err != nil {
			print(err.Error())
			os.Exit(1)
		}
		fmt.Printf("Pid: %d\n", cmd.Process.Pid)
		ch <- cmd.Wait()
	}()

	time.Sleep(2 * time.Second)
	fmt.Printf("%v\n", cmd.Process.Signal(syscall.SIGSTOP))
	time.Sleep(2 * time.Second)
	fmt.Printf("%v\n", cmd.Process.Signal(syscall.SIGCONT))
	//fmt.Printf("%v\n", syscall.Kill(cmd.Process.Pid, syscall.SIGCONT))
}
