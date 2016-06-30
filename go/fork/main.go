package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func waitForExit() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
	os.Exit(0)
}

func main() {
	pid := os.Getpid()
	parentpid := os.Getppid()
	fmt.Printf("Parent: %d pid: %d ", pid, parentpid)

	proc, err := os.FindProcess(pid)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Process: %d\n", proc.Pid)

	go func() {
		r, w, err := os.Pipe()
		if err != nil {
			panic(err)
		}
		defer r.Close()
		cmdToRun := "/bin/sleep"
		args := []string{"name-of-the-cmd", "30"}
		procAttr := new(os.ProcAttr)
		procAttr.Files = []*os.File{nil, w, os.Stderr}
		process, err := os.StartProcess(cmdToRun, args, procAttr)
		if err != nil {
			fmt.Printf("ERROR Unable to run %s: %s\n", cmdToRun, err.Error())
			return
		}
		fmt.Printf("%s running as pid %d\n", cmdToRun, process.Pid)

		processState, err := process.Wait()

		if err != nil {
			panic(err)
		}

		err = process.Release()

		if err != nil {
			panic(err)
		}

		fmt.Println("Did the child process exited? : ", processState.Exited())
		fmt.Println("So the child pid is? : ", processState.Pid())
		fmt.Println("Exited successfully? : ", processState.Success())

		fmt.Println("Exited system CPU time ? : ", processState.SystemTime())
		fmt.Println("Exited user CPU time ? : ", processState.UserTime())

		// just to be sure, let's kill again
		err = process.Signal(syscall.SIGKILL)

		if err != nil {
			fmt.Println(err) // see what the serial killer has to say
			return
		}
		w.Close()
	}()

	waitForExit()
}
