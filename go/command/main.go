package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os/exec"
)

func main() {
	cmd := exec.Command("/tmp/stdout")
	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)

	cmd.Stdout = stdout
	cmd.Stderr = stderr

	if err := cmd.Start(); err != nil {
		panic(err)
	}

	if err := cmd.Wait(); err != nil {
		panic(err)
	}

	in := bufio.NewScanner(io.MultiReader(stdout, stderr))
	for in.Scan() {
		fmt.Println(in.Text())
	}

}
