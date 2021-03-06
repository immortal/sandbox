package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"syscall"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s pid", os.Args[0])
		os.Exit(1)
	}

	pid, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		panic(err)
	}

	process, _ := os.FindProcess(int(pid))

	kq, err := syscall.Kqueue()
	if err != nil {
		fmt.Println(err)
	}

	ev1 := syscall.Kevent_t{
		Ident:  uint64(process.Pid),
		Filter: syscall.EVFILT_PROC,
		Flags:  syscall.EV_ADD,
		Fflags: syscall.NOTE_EXIT,
		Data:   0,
		Udata:  nil,
	}

	for {
		events := make([]syscall.Kevent_t, 1)
		n, err := syscall.Kevent(kq, []syscall.Kevent_t{ev1}, events, nil)
		if err != nil {
			log.Println("Error creating kevent")
		}
		if n > 0 {
			break
		}
	}

	fmt.Println("fin")
}
