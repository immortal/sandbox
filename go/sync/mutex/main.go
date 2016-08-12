package main

import (
	"fmt"
	"os/exec"
	"sync"
	"time"
)

type Daemon struct {
	sync.Mutex
	cmd   string
	state uint32
}

func (self *Daemon) Run(i int) error {
	self.Lock()
	defer self.Unlock()
	println(i)
	cmd := exec.Command("sleep", "2")
	if err := cmd.Start(); err != nil {
		return err
	}
	pid := cmd.Process.Pid

	fmt.Printf("cmd: %d\n", pid)
	return cmd.Wait()
}

func main() {
	d := &Daemon{cmd: "sleep"}

	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		go func(x int) {
			if err := d.Run(x); err != nil {
				fmt.Println(err)
			}
		}(i)
	}

	for {
	}
}
