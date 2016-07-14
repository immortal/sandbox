package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"syscall"
	"time"
)

func sleep(ch chan<- bool) {
	go func() {
		time.Sleep(1 * time.Second)
		ch <- true
	}()
}

func log(reader *io.PipeReader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	reader.Close()
}

func main() {
	exit := make(chan error, 1)
	go run(exit)

	for {
		select {
		case err := <-exit:
			if err != nil {
				println(err.Error())
				if err.Error() == "EXIT" {
					println("PID exited")
				}
			}
			println("fin, restarting")
			//run(exit)
		}
	}
}

func run(ch chan<- error) {
	//cmd := exec.Command("/tmp/stdout")
	cmd := exec.Command("bundle", "exec", "unicorn", "-c", "unicorn.rb")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
		Pgid:    0,
	}

	r, w := io.Pipe()
	cmd.Stdout = w
	cmd.Stderr = w

	go log(r)

	go func() {
		// Close the writer or the pipe will not be closed for c2
		defer w.Close()

		if err := cmd.Start(); err != nil {
			print(err.Error())
			os.Exit(1)
		}
		fmt.Printf("Pid: %d\n", cmd.Process.Pid)
		ch <- cmd.Wait()
	}()
	go func() {
		pid, err := readPidfile("./unicorn.pid")
		if err != nil {
			println(err.Error())
			return
		}
		println(pid, "<------------")
		watchPid(pid, ch)
	}()
}
