package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	init := "./client"
	initArgs := []string{"hello world"}

	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	cmd := exec.Command(init, initArgs...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.ExtraFiles = []*os.File{w}

	if err := cmd.Start(); err != nil {
		panic(err)
	}
	var data interface{}
	decoder := json.NewDecoder(r)
	if err := decoder.Decode(&data); err != nil {
		panic(err)
	}
	fmt.Printf("Data received from child pipe: %v\n", data)
}
