package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}
	arg := strings.ToUpper(os.Args[1])

	pipe := os.NewFile(uintptr(3), "pipe")
	err := json.NewEncoder(pipe).Encode(arg)
	if err != nil {
		panic(err)
	}
	fmt.Println("This message printed to standard output, not to the pipe")
}
