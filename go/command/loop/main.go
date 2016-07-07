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

func run(c []string) error {
	cmd := exec.Command(c[1], c[2:]...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}
	if err := cmd.Start(); err != nil {
		return err
	}
	go out(stdout)
	go out(stderr)

	return cmd.Wait()
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Enter a command")
		os.Exit(1)
	}

	for {
		err := run(os.Args)
		if err != nil {
			println(err)
		}
	}
}
