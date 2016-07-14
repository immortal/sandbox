// +build freebsd netbsd openbsd dragonfly darwin

package main

import (
	"errors"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func readPidfile(pid_file string) (int, error) {
	time.Sleep(3 * time.Second)
	content, err := ioutil.ReadFile(pid_file)
	if err != nil {
		return 0, err
	}
	lines := strings.Split(string(content), "\n")
	pid, err := strconv.Atoi(lines[0])
	if err != nil {
		return 0, err
	}
	return pid, nil
}

func watchPid(pid int, ch chan<- error) {
	kq, err := syscall.Kqueue()
	if err != nil {
		panic(err)
	}

	ev1 := syscall.Kevent_t{
		Ident:  uint64(pid),
		Filter: syscall.EVFILT_PROC,
		Flags:  syscall.EV_ADD | syscall.EV_ENABLE | syscall.EV_ONESHOT,
		Fflags: syscall.NOTE_EXIT,
		Data:   0,
		Udata:  nil,
	}

	for {
		events := make([]syscall.Kevent_t, 1)
		n, err := syscall.Kevent(kq, []syscall.Kevent_t{ev1}, events, nil)
		if err != nil {
			println(os.NewSyscallError("kevent", err))
			return
		}
		for i := 0; i < n; i++ {
			syscall.Close(kq)
			ch <- errors.New("EXIT")
			return
		}
	}
}
