package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s /path", os.Args[0])
		os.Exit(1)
	}

	fd, err := syscall.Open(os.Args[1], syscall.O_RDONLY, 0)
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
		Flags:  syscall.EV_ADD | syscall.EV_ENABLE | syscall.EV_ONESHOT,
		Fflags: syscall.NOTE_DELETE | syscall.NOTE_WRITE | syscall.NOTE_EXTEND | syscall.NOTE_ATTRIB | syscall.NOTE_LINK | syscall.NOTE_RENAME | syscall.NOTE_REVOKE,
		Data:   0,
		Udata:  nil,
	}

Loop:
	// wait for events
	for {
		// create kevent
		events := make([]syscall.Kevent_t, 10)
		n, err := syscall.Kevent(kq, []syscall.Kevent_t{ev1}, events, nil)
		if err != nil {
			log.Println("Error creating kevent")
		}
		// check if there was an event
		for i := 0; i < n; i++ {
			log.Printf("Event [%d] -> %+v data: %#v", i, events[i], events[i].Data)

			if events[i].Fflags == syscall.NOTE_RENAME {
				print("renamed")
			}
			// Fflags:17
			if events[i].Fflags == syscall.NOTE_DELETE || events[i].Fflags == 17 {
				print("deleted")
				break Loop
			}
		}
	}

	fmt.Println("fin")
}
