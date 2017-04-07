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

	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	kq, err := syscall.Kqueue()
	if err != nil {
		fmt.Println(err)
	}

	ev1 := syscall.Kevent_t{
		Ident:  uint64(file.Fd()),
		Filter: syscall.EVFILT_VNODE,
		Flags:  syscall.EV_ADD | syscall.EV_ENABLE | syscall.EV_ONESHOT,
		Fflags: syscall.NOTE_DELETE | syscall.NOTE_WRITE | syscall.NOTE_EXTEND | syscall.NOTE_ATTRIB | syscall.NOTE_LINK | syscall.NOTE_RENAME | syscall.NOTE_REVOKE,
		Data:   0,
		Udata:  nil,
	}

	//Loop:
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
				//				break Loop
			}
			if n > 0 {
				fmt.Printf("n = %+v\n", n)
				return
			}
		}
	}

	fmt.Println("fin")
}
