package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Printf("usage: %s /path/file", os.Args[0])
		os.Exit(1)
	}

	//	syscall.Mknod(os.Args[1], syscall.S_IFIFO|0666, 0)
	file, err := os.OpenFile(os.Args[1], os.O_RDONLY, os.ModeNamedPipe)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)

	line, err := reader.ReadString(1)

	for err != io.EOF {
		fmt.Print(line)
		line, err = reader.ReadString(1)
	}
}
