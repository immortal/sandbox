package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"syscall"
)

func WatchDir(dir string, ch chan<- string) {
	file, err := os.Open(dir)
	if err != nil {
		log.Printf("err = %+v\n", err)
	}

	kq, err := syscall.Kqueue()
	if err != nil {
		log.Printf("err = %+v\n", err)
	}

	ev1 := syscall.Kevent_t{
		Ident:  uint64(file.Fd()),
		Filter: syscall.EVFILT_VNODE,
		Flags:  syscall.EV_ADD | syscall.EV_ENABLE | syscall.EV_ONESHOT,
		Fflags: syscall.NOTE_DELETE | syscall.NOTE_WRITE | syscall.NOTE_EXTEND | syscall.NOTE_ATTRIB | syscall.NOTE_LINK | syscall.NOTE_RENAME | syscall.NOTE_REVOKE,
		Data:   0,
		Udata:  nil,
	}

	// create kevent
	events := make([]syscall.Kevent_t, 1)
	n, err := syscall.Kevent(kq, []syscall.Kevent_t{ev1}, events, nil)
	if err != nil {
		log.Println("Error creating kevent")
	}
	// check if there was an event
	for {
		if n > 0 {
			ch <- dir
			return
		}
	}
}

func Scandir(dir string) []string {
	yml := []string{}
	d, err := os.Open(dir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer d.Close()

	files, err := d.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, file := range files {
		if file.Mode().IsRegular() {
			if filepath.Ext(file.Name()) == ".yml" {
				yml = append(yml, file.Name())
			}
		}
	}
	return yml
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s /path", os.Args[0])
		os.Exit(1)
	}

	dir := os.Args[1]

	watchDir := make(chan string, 1)
	watchFile := make(chan string, 1)

	yml := Scandir(dir)
	for _, y := range yml {
		fmt.Printf("Watching  %s\n", y)
		go WatchDir(filepath.Join(dir, y), watchFile)
	}

	WatchDir(dir, watchDir)

	for {
		select {
		case dir := <-watchDir:
			fmt.Printf("dir = %s\n", dir)
			println("find *.yml")
			yml2 := Scandir(dir)
			// replace this with a map On2
			for _, y := range yml2 {
				var skip bool
				for _, oy := range yml {
					if oy == y {
						skip = true
						continue
					}
				}
				if !skip {
					fmt.Printf("Watching  %s\n", y)
					go WatchDir(filepath.Join(dir, y), watchFile)
				}
			}
			go WatchDir(dir, watchDir)
		case file := <-watchFile:
			fmt.Printf("file changed = %s\n", file)
			go WatchDir(file, watchFile)
		}
	}
}
