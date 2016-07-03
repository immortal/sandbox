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

	// kevent
	// struct kevent {
	//    uintptr_t ident;		 /*	identifier for this event */
	//    short     filter;		 /*	filter for event */
	//    u_short   flags;		 /*	action flags for kqueue	*/
	//    u_int     fflags;		 /*	filter flag value */
	//    intptr_t  data;		 /*	filter data value */
	//    void      *udata;		 /*	opaque user data identifier */
	// };
	//
	// pwait.c (FreeBSD)
	// EV_SET(e + nleft, pid, EVFILT_PROC, EV_ADD, NOTE_EXIT, 0, NULL);
	ev1 := syscall.Kevent_t{
		Ident:  uint64(process.Pid),
		Filter: syscall.EVFILT_PROC,
		Flags:  syscall.EV_ADD,
		Fflags: syscall.NOTE_EXIT,
		Data:   0,
		Udata:  nil,
	}

	// configure timeout
	//	timeout := syscall.Timespec{
	//		Sec:  1,
	//		Nsec: 0,
	//	}

	// wait for events
	for {
		// create kevent
		events := make([]syscall.Kevent_t, 1)
		// check https://golang.org/src/syscall/syscall_bsd.go
		//
		// func Kevent(kq int, changes, events []Kevent_t, timeout *Timespec) (n int, err error)
		//
		// n, err := syscall.Kevent(kq, []syscall.Kevent_t{ev1}, events, &timeout)
		n, err := syscall.Kevent(kq, []syscall.Kevent_t{ev1}, events, nil)
		if err != nil {
			log.Println("Error creating kevent")
		}
		// check if there was an event
		for i := 0; i < n; i++ {
			log.Printf("Event [%d] -> %+v data: %#v", i, events[i], events[i].Data)
		}
		if n > 0 {
			break
		}
	}

	fmt.Println("fin")
}
