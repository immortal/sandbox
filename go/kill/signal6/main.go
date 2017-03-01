package main

import (
	"fmt"
	"log"
	"os/exec"
	"syscall"
	"unsafe"
)

func main() {
	cmd := exec.Command("sleep", "10")
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	// signal when wait4 will return immediately
	go func() {
		var siginfo [128]byte
		psig := &siginfo[0]
		_, _, e := syscall.Syscall6(syscall.SYS_WAITID, 1, uintptr(cmd.Process.Pid), uintptr(unsafe.Pointer(psig)), syscall.WEXITED|syscall.WNOWAIT, 0, 0)
		fmt.Println("WAITID RETURNED -- this shouldn't happen:", e)
	}()

	err := cmd.Process.Signal(syscall.SIGSTOP)
	if err != nil {
		log.Fatal(err)
	}
	cmd.Wait()
}
