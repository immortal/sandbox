package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Printf("usage: %s /path/file", os.Args[0])
		os.Exit(1)
	}

	//	syscall.Mknod(os.Args[1], syscall.S_IFIFO|0666, 0)
	file, err := os.OpenFile(os.Args[1], os.O_RDONLY, os.ModeNamedPipe)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
