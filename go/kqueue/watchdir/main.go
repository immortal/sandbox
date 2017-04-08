package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"syscall"
	"time"
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
			file.Close()
			ch <- dir
			return
		}
	}
}

func Scandir(dir string) map[string]string {
	yml := map[string]string{}
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
				yml[file.Name()] = filepath.Join(dir, file.Name())
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
	for f, p := range yml {
		fmt.Printf("FIRST -- Watching  %s in path %s\n", f, p)
		go WatchDir(p, watchFile)
	}

	go WatchDir(dir, watchDir)

	for {
		println("loop..")
		select {
		case dir := <-watchDir:
			fmt.Printf("dir = %s\n", dir)
			println("find *.yml")
			newFiles := Scandir(dir)
			// replace this with a map On2
			for f, p := range newFiles {
				if _, ok := yml[f]; !ok {
					fmt.Printf("Watching new file %s in path %s\n", f, p)
					yml[f] = p
					go WatchDir(p, watchFile)
				}
			}
			fmt.Printf("Nothing to add, watching again dir = %s\n", dir)
			go WatchDir(dir, watchDir)
		case file := <-watchFile:
			fmt.Printf("file changed = %s\n", file)
			go WatchDir(file, watchFile)
		}
		// to avoid err = too many open files,  Error creating kevent
		time.Sleep(100 * time.Millisecond)
	}
}
