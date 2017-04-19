package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
)

const (
	modifyFlags = syscall.NOTE_WRITE | syscall.NOTE_EXTEND | syscall.NOTE_ATTRIB
	deleteFlags = syscall.NOTE_DELETE | syscall.NOTE_RENAME | syscall.NOTE_REVOKE
	allFlags    = modifyFlags | deleteFlags
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s /path", os.Args[0])
		os.Exit(1)
	}

	fd, err := syscall.Open(os.Args[1], syscall.O_EVTONLY, 0700)
	if err != nil {
		panic(err)
	}

	kq, err := syscall.Kqueue()
	if err != nil {
		fmt.Println(err)
	}

	ev1 := syscall.Kevent_t{
		Ident:  uint64(fd),
		Filter: syscall.EVFILT_VNODE,
		Flags:  syscall.EV_ADD | syscall.EV_ENABLE | syscall.EV_ONESHOT | syscall.EV_CLEAR,
		Fflags: syscall.NOTE_DELETE | syscall.NOTE_WRITE | syscall.NOTE_EXTEND | syscall.NOTE_ATTRIB | syscall.NOTE_LINK | syscall.NOTE_RENAME | syscall.NOTE_REVOKE,
		Data:   0,
		Udata:  nil,
	}

	//Loop:
	// wait for events
	for {
		println("Creating event")
		// create kevent
		events := []syscall.Kevent_t{ev1}
		n, err := syscall.Kevent(kq, events, events, nil)
		if err != nil {
			log.Println("Error creating kevent")
		}
		if n != 1 {
			fmt.Printf("n = %+v\n", n)
			return
		}
		// check if there was an event
		for _, ev := range events {
			if ev.Fflags&modifyFlags != 0 {
				println("modified")
			} else if ev.Fflags&deleteFlags != 0 {
				println("deleted")
			} else {
				fmt.Printf("didn't catch ev.Flags = %+v\n", ev.Flags)
			}
		}
	}

	fmt.Println("fin")
}
