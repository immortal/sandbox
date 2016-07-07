package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"
)

func Log(s interface{}) {
	t := time.Now().UTC().Format(time.RFC3339Nano)
	log := fmt.Sprintf("%s %v\n", t, s)
	f, err := os.OpenFile("/tmp/test.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err = f.WriteString(log); err != nil {
		panic(err)
	}
}

func out(p io.ReadCloser) {
	in := bufio.NewScanner(p)
	for in.Scan() {
		Log(in.Text())
	}
}

func run(c []string, ch chan<- error) {
	cmd := exec.Command(c[1], c[2:]...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		ch <- err
		return
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		ch <- err
		return
	}
	if err := cmd.Start(); err != nil {
		ch <- err
		return
	}
	go out(stdout)
	go out(stderr)

	ch <- cmd.Wait()
}

func main() {

	status := make(chan error, 1)
	if len(os.Args) < 2 {
		fmt.Println("Enter a command")
		os.Exit(1)
	}

	run(os.Args, status)

	for {
		select {
		case err := <-status:
			if err != nil {
				fmt.Printf("Status: %#v\n", err.Error())
			}
			time.Sleep(1 * time.Second)
			run(os.Args, status)
		}
	}
}
